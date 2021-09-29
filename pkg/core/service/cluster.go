package service

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
	m "github.com/cloud-barista/cb-mcas/pkg/core/model/mcks"
)

func ListCluster(namespace string) (*model.ClusterList, error) {
	clusterList := model.NewClusterList(namespace)

	mcks := m.NewMcks(namespace)

	mcksClusterList, err := mcks.ListCluster()
	if err != nil {
		return nil, err
	}

	for _, mcksCluster := range mcksClusterList.Items {
		cluster := model.NewCluster(mcksCluster.Namespace, mcksCluster.Name)
		cluster.SetMcis(mcksCluster.Mcis)
		for _, mcksNode := range mcksCluster.Nodes {
			node := model.NewNode(
				mcksNode.Name,
				mcksNode.PublicIp,
				mcksNode.Csp,
				mcksNode.Role,
				mcksNode.Spec,
			)
			cluster.AddNode(node)
		}
		cluster.SetStatus(mcksCluster.Status)

		clusterList.Items = append(clusterList.Items, *cluster)
	}

	return clusterList, nil
}

func GetCluster(namespace, name string) (*model.Cluster, error) {
	mcks := m.NewMcks(namespace)
	mcksCluster, err := mcks.GetCluster(name)
	if err != nil {
		return nil, err
	}

	cluster := model.NewCluster(mcksCluster.Namespace, mcksCluster.Name)
	cluster.SetMcis(mcksCluster.Mcis)

	for _, mcksNode := range mcksCluster.Nodes {
		node := model.NewNode(
			mcksNode.Name,
			mcksNode.PublicIp,
			mcksNode.Csp,
			mcksNode.Role,
			mcksNode.Spec,
		)
		cluster.AddNode(node)
	}

	cluster.SetStatus(mcksCluster.Status)

	return cluster, nil
}

func CreateCluster(namespace string, req *model.ClusterReq) (*model.Cluster, error) {
	mcks := m.NewMcks(namespace)
	mcksCluster, err := mcks.CreateCluster(req.McksClusterReq)
	if err != nil {
		return nil, err
	}

	cluster := model.NewCluster(mcksCluster.Namespace, mcksCluster.Name)
	cluster.SetMcis(mcksCluster.Mcis)

	for _, mcksNode := range mcksCluster.Nodes {
		node := model.NewNode(
			mcksNode.Name,
			mcksNode.PublicIp,
			mcksNode.Csp,
			mcksNode.Role,
			mcksNode.Spec,
		)
		cluster.AddNode(node)
	}

	cluster.SetStatus(mcksCluster.Status)

	return cluster, nil
}

func DeleteCluster(namespace, name string) (*model.Status, error) {
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

	return status, nil
}
