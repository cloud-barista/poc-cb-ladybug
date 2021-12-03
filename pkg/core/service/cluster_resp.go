package service

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/cb-ladybug/pkg/core/model"
	m "github.com/cloud-barista/cb-ladybug/pkg/core/model/mcks"
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

	var clusterResp *model.ClusterResp = nil
	if mcksCluster != nil {
		clusterResp = makeClusterResp(mcksCluster)
	}

	return clusterResp, nil
}

func CreateCluster(namespace string, req *model.ClusterReq) (*model.ClusterResp, error) {
	mcks := m.NewMcks(namespace)
	mcksCluster, err := mcks.CreateCluster(req.McksClusterReq)
	if err != nil {
		return nil, err
	}

	clusterResp := makeClusterResp(mcksCluster)

	return clusterResp, nil
}

func DeleteCluster(namespace, name string) (*model.Status, error) {
	status := model.NewStatus()

	common.CBLog.Infof("delete the cluster '%s'", name)

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

func makeClusterResp(mcksCluster *m.McksCluster) *model.ClusterResp {
	clusterResp := model.NewClusterResp(mcksCluster.Namespace, mcksCluster.Name)
	clusterResp.SetCreatedTime(mcksCluster.CreatedTime)
	clusterResp.SetDescription(mcksCluster.Description)
	clusterResp.SetLabel(mcksCluster.Label)
	clusterResp.SetMcis(mcksCluster.Mcis)

	for _, mcksNode := range mcksCluster.Nodes {
		node := model.NewNode(
			mcksNode.CreatedTime,
			mcksNode.Csp,
			mcksNode.CspLabel,
			mcksNode.Name,
			mcksNode.PublicIp,
			mcksNode.RegionLabel,
			mcksNode.Role,
			mcksNode.Spec,
			mcksNode.ZoneLabel,
		)
		clusterResp.AddNode(node)
	}

	clusterResp.SetStatus(mcksCluster.Status.Phase, mcksCluster.Status.Reason, mcksCluster.Status.Message)
	clusterResp.SetClusterConfig(mcksCluster.ClusterConfig)

	return clusterResp
}
