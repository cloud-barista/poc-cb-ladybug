package service

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
	m "github.com/cloud-barista/cb-mcas/pkg/core/model/mcks"
)

func ListCluster(namespace string) (*model.ClusterRespList, error) {
	clusterRespList := model.NewClusterRespList(namespace)

	mcks := m.NewMcks(namespace)

	mcksClusterList, err := mcks.ListCluster()
	if err != nil {
		return nil, err
	}

	for _, mcksCluster := range mcksClusterList.Items {
		clusterResp := makeClusterResp(&mcksCluster)
		clusterRespList.Items = append(clusterRespList.Items, *clusterResp)
	}

	return clusterRespList, nil
}

func GetCluster(namespace, name string) (*model.ClusterResp, error) {
	mcks := m.NewMcks(namespace)
	mcksCluster, err := mcks.GetCluster(name)
	if err != nil {
		return nil, err
	}

	clusterResp := makeClusterResp(mcksCluster)

	return clusterResp, nil
}

func CreateCluster(namespace string, req *model.ClusterReq) (*model.ClusterResp, error) {
	/*
		clusterInfo := model.NewClusterInfo(namespace, req.Name)
		clusterInfo.TriggerCreate()
	*/
	mcks := m.NewMcks(namespace)
	mcksCluster, err := mcks.CreateCluster(req.McksClusterReq)
	if err != nil {
		return nil, err
	}

	clusterResp := makeClusterResp(mcksCluster)

	/*
		if mcksCluster.Status == m.MCKS_CLUSTER_STATUS_COMPLETED {
			clusterInfo.TriggerSuccess()
		} else {
			clusterInfo.TriggerFail()
		}
	*/

	return clusterResp, nil
}

func DeleteCluster(namespace, name string) (*model.Status, error) {
	/*
		clusterInfo := model.NewClusterInfo(namespace, name)
		ci_err := clusterInfo.Load()
		if ci_err != nil {
			common.CBLog.Warnf("cluster '%s' not found in store", name)
		}
	*/
	status := model.NewStatus()

	common.CBLog.Infof("delete a cluster (name=%s)", name)

	mcks := m.NewMcks(namespace)
	mcksStatus, err := mcks.DeleteCluster(name)
	if err != nil {
		return nil, err
	}

	common.CBLog.Infof("MCKS return Status{Code=%d, Message=%s}", mcksStatus.Code, mcksStatus.Message)

	if mcksStatus.Code == m.MCKS_STATUS_SUCCESS {
		status.Code = model.STATUS_SUCCESS
		if strings.Contains(mcksStatus.Message, "not found") {
			common.CBLog.Infof("cluster '%s' not found", name)
			status.Message = fmt.Sprintf("cluster '%s' not found", name)
		} else {
			common.CBLog.Infof("cluster '%s' has been deleted", name)
			status.Message = fmt.Sprintf("cluster '%s' has been deleted", name)
		}
	} else if mcksStatus.Code == m.MCKS_STATUS_UNKNOWN {
		common.CBLog.Infof("unknown error occurred when deleting the cluster '%s'", name)
		status.Code = model.STATUS_FAIL
		status.Message = fmt.Sprintf("unknown error occurred when deleting the cluster '%s'", name)
	} else { // mcksStatus.Code == mcks.MCKS_STATUS_NOT_EXIST
		common.CBLog.Infof("cluster '%s' was not found", name)
		status.Code = model.STATUS_FAIL
		status.Message = fmt.Sprintf("cluster '%s' was not found", name)
	}

	/*
		if ci_err != nil {
			clusterInfo.TriggerDelete()
			clusterInfo.Delete()
		}
	*/
	return status, nil
}

func makeClusterResp(mcksCluster *m.McksCluster) *model.ClusterResp {
	/*
		clusterInfo := model.NewClusterInfo(mcksCluster.Namespace, mcksCluster.Name)
		ci_err := clusterInfo.Load()
		if ci_err != nil {
			common.CBLog.Warnf("cluster '%s' not found in store", mcksCluster.Name)
			clusterInfo.TriggerCreate()
		}
	*/
	clusterResp := model.NewClusterResp(mcksCluster.Namespace, mcksCluster.Name)
	clusterResp.SetMcis(mcksCluster.Mcis)

	for _, mcksNode := range mcksCluster.Nodes {
		node := model.NewNode(
			mcksNode.Name,
			mcksNode.PublicIp,
			mcksNode.Csp,
			mcksNode.Role,
			mcksNode.Spec,
		)
		clusterResp.AddNode(node)
	}

	clusterResp.SetStatus(mcksCluster.Status)

	/*
		if mcksCluster.Status == m.MCKS_CLUSTER_STATUS_COMPLETED {
			clusterInfo.SetClusterConfig(mcksCluster.ClusterConfig)
			clusterInfo.TriggerSuccess()
		} else {
			if ci_err != nil {
				clusterInfo.TriggerFail()
			} else {
				clusterInfo.TriggerUnknown()
			}
		}
	*/
	return clusterResp
}
