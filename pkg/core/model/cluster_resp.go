package model

import (
	"errors"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/model/mcks"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/config"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/lang"
)

type (
	Node struct {
		CreatedTime string `json:"createdTime"`
		Csp         string `json:"csp"`
		CspLabel    string `json:"cspLabel"`
		Name        string `json:"name"`
		PublicIp    string `json:"publicIp"`
		RegionLabel string `json:"regionLabel"`
		Role        string `json:"role"`
		Spec        string `json:"spec"`
		ZoneLabel   string `json:"zoneLabel"`
	}

	ClusterResp struct {
		Model
		CreatedTime string `json:"createdTime"`
		Description string `json:"description"`
		Label       string `json:"label"`
		Namespace   string `json:"namespace"`
		Mcis        string `json:"mcis"`
		Nodes       []Node `json:"nodes"`
		Status      struct {
			Phase   string `json:"phase"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"status"`

		clusterConfig string
	}

	ClusterReq struct {
		mcks.McksClusterReq
	}

	ClusterRespList struct {
		ListModel
		Namespace string        `json:"namespace"`
		Items     []ClusterResp `json:"items"`
	}
)

func NewNode(createdTime, csp, cspLabel, name, publicIp, regionLabel, role, spec, zoneLabel string) *Node {
	return &Node{
		CreatedTime: createdTime,
		Csp:         csp,
		CspLabel:    cspLabel,
		Name:        name,
		PublicIp:    publicIp,
		RegionLabel: regionLabel,
		Role:        role,
		Spec:        spec,
		ZoneLabel:   zoneLabel,
	}
}

func NewClusterResp(namespace, name string) *ClusterResp {
	return &ClusterResp{
		Model:     Model{Kind: KIND_CLUSTER_RESP, Name: name},
		Namespace: namespace,
	}
}

func (self *ClusterResp) SetCreatedTime(createdTime string) {
	self.CreatedTime = createdTime
}

func (self *ClusterResp) SetDescription(description string) {
	self.Description = description
}

func (self *ClusterResp) SetLabel(label string) {
	self.Label = label
}

func (self *ClusterResp) SetMcis(mcis string) {
	self.Mcis = mcis
}

func (self *ClusterResp) AddNode(node *Node) {
	self.Nodes = append(self.Nodes, *node)
}

func (self *ClusterResp) SetStatus(phase, reason, message string) {
	self.Status.Phase = phase
	self.Status.Reason = reason
	self.Status.Message = message
}

func (self *ClusterResp) GetClusterConfig() string {
	return self.clusterConfig
}

func (self *ClusterResp) SetClusterConfig(clusterConfig string) {
	self.clusterConfig = clusterConfig
}

func NewClusterRespList(namespace string) *ClusterRespList {
	return &ClusterRespList{
		ListModel: ListModel{Kind: KIND_CLUSTER_RESP_LIST},
		Namespace: namespace,
		Items:     []ClusterResp{},
	}
}

func ClusterReqConfKubeDef(req *ClusterReq) {
	req.Config.Kubernetes.NetworkCni = lang.NVL(req.Config.Kubernetes.NetworkCni, config.NETWORKCNI_KILO)
	req.Config.Kubernetes.PodCidr = lang.NVL(req.Config.Kubernetes.PodCidr, config.POD_CIDR)
	req.Config.Kubernetes.ServiceCidr = lang.NVL(req.Config.Kubernetes.ServiceCidr, config.SERVICE_CIDR)
	req.Config.Kubernetes.ServiceDnsDomain = lang.NVL(req.Config.Kubernetes.ServiceDnsDomain, config.SERVICE_DOMAIN)
}

func ClusterReqValidate(req *ClusterReq) error {
	if len(req.ControlPlane) == 0 {
		return errors.New("control plane node must be at least one")
	}
	if len(req.ControlPlane) > 1 {
		return errors.New("only one control plane node is supported")
	}
	if len(req.Worker) == 0 {
		return errors.New("worker node must be at least one")
	}
	if !(req.Config.Kubernetes.NetworkCni == config.NETWORKCNI_CANAL || req.Config.Kubernetes.NetworkCni == config.NETWORKCNI_KILO) {
		return errors.New("network cni allows only kilo or canal")
	}

	if len(req.Name) == 0 {
		return errors.New("cluster name is empty")
	} else {
		err := lang.CheckName(req.Name)
		if err != nil {
			return err
		}
	}

	if len(req.Config.Kubernetes.PodCidr) > 0 {
		err := lang.CheckIpCidr("podCidr", req.Config.Kubernetes.PodCidr)
		if err != nil {
			return err
		}
	}
	if len(req.Config.Kubernetes.ServiceCidr) > 0 {
		err := lang.CheckIpCidr("serviceCidr", req.Config.Kubernetes.ServiceCidr)
		if err != nil {
			return err
		}
	}

	return nil
}
