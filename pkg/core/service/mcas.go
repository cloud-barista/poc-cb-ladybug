package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/cb-ladybug/pkg/core/model"
	m "github.com/cloud-barista/cb-ladybug/pkg/core/model/mcks"
)

func GetMcas(namespace string) (string, error) {
	mcas := model.NewMcas(namespace)
	status, err := mcas.GetStatus()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err2 := mcas.Disable()
			if err2 != nil {
				common.CBLog.Error(err2)
				return "", err2
			}
			status, err2 = mcas.GetStatus()
			if err2 != nil {
				common.CBLog.Error(err2)
				return "", err2
			}
		} else {
			common.CBLog.Error(err)
			return "", err
		}
	}

	return string(status), nil
}

func EnableMcas(namespace string) error {
	mcas := model.NewMcas(namespace)
	/*
		status, err := mcas.GetStatus()
		if err != nil {
			common.CBLog.Error(err)
			return err
		}

		if status == model.STATUS_MCAS_ENABLED {
			common.CBLog.Infof("MCAS for namespace '%s' is already enabled.\n", namespace)
			return nil
		}
	*/

	clusterName := "mcas-cluster"

	clusterResp, err := GetCluster(namespace, clusterName)
	if err != nil {
		return err
	}

	if clusterResp != nil && clusterResp.Status.Phase == m.MCKS_CLUSTER_STATUS_PHASE_PROVISIONED {
		common.CBLog.Infof("'%s/%s' cluster is already existed.\n", namespace, clusterName)
		return nil
	}

	//
	// FIXME: fix the repositoy url
	//
	addPackageRepo(namespace, "http://localhost:38080")

	//
	// Create a new cluster
	//

	clusterInfo := model.NewClusterInfo(namespace, clusterName)
	clusterInfo.TriggerCreate()

	clusterReq := makeClusterReq(clusterName)
	clusterResp, err = CreateCluster(namespace, clusterReq)
	if err != nil {
		common.CBLog.Error(err)
		//common.CBLog.Infof("try to delete the cluster '%s/%s'", namespace, clusterName)
		//DeleteCluster(namespace, clusterName)
		return err
	}

	if clusterResp.Status.Phase != m.MCKS_CLUSTER_STATUS_PHASE_PROVISIONED {
		clusterInfo.TriggerFail()
		return errors.New(fmt.Sprintf("cluster '%s/%s' creating is failed (reason=%s, message=%s)",
			namespace, clusterName, clusterResp.Status.Reason, clusterResp.Status.Message))
	} else {
		clusterInfo.SetClusterConfig(clusterResp.GetClusterConfig())
		clusterInfo.TriggerSuccess()
	}

	//
	// Set MCAS Status to 'enabled'
	//
	err = mcas.Enable()
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	return nil
}

func DisableMcas(namespace string) error {
	mcas := model.NewMcas(namespace)

	/*
		status, err := mcas.GetStatus()
		if err != nil {
			common.CBLog.Error(err)
			return err
		}

		if status == model.STATUS_MCAS_DISABLED {
			common.CBLog.Infof("MCAS for namespace '%s' is already disabled.\n", namespace)
			return nil
		}
	*/

	//
	// Delete the mcas cluster
	//
	clusterName := "mcas-cluster"

	clusterInfo := model.NewClusterInfo(namespace, clusterName)
	ci_err := clusterInfo.Select()
	if ci_err == nil {
		clusterInfo.TriggerDelete()
	}

	clusterStatus, err := DeleteCluster(namespace, clusterName)
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	if ci_err == nil {
		clusterInfo.Delete()
	}

	_ = clusterStatus

	//
	// Set MCAS Status to 'disabled'
	//
	err = mcas.Disable()
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	return nil
}

func makeClusterReq(clusterName string) *model.ClusterReq {
	var clusterReq model.ClusterReq

	clusterReq.Config.Kubernetes.NetworkCni = "kilo"
	clusterReq.Config.Kubernetes.PodCidr = "10.244.0.0/16"
	clusterReq.Config.Kubernetes.ServiceCidr = "10.96.0.0/12"
	clusterReq.Config.Kubernetes.ServiceDnsDomain = "cluster.local"

	var ncCp m.McksNodeConfig
	ncCp.Connection = "config-aws-ap-northeast-2"
	ncCp.Count = 1
	ncCp.Spec = "t2.medium"

	clusterReq.ControlPlane = append(clusterReq.ControlPlane, ncCp)
	clusterReq.Description = "cluster for MCAS"
	clusterReq.InstallMonAgent = "no"
	clusterReq.Label = "MCAS"

	clusterReq.Name = clusterName

	var ncW m.McksNodeConfig
	ncW.Connection = "config-aws-ap-northeast-1"
	ncW.Count = 1
	ncW.Spec = "t2.small"

	clusterReq.Worker = append(clusterReq.Worker, ncW)

	ncW.Connection = "config-gcp-asia-northeast3"
	ncW.Count = 1
	ncW.Spec = "n1-standard-2"

	clusterReq.Worker = append(clusterReq.Worker, ncW)

	return &clusterReq
}
