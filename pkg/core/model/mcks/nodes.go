package mcks

import (
	"fmt"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type (
	McksNode struct {
		Credential string `json:"credential"`
		Csp        string `json:"csp"`
		Kind       string `json:"kind"`
		Name       string `json:"name"`
		PublicIp   string `json:"publicIp"`
		Role       string `json:"role"`
		Spec       string `json:"spec"`
		Uid        string `json:"uid"`
	}

	McksNodeList struct {
		Kind  string     `json:"kind"`
		Items []McksNode `json:"items"`
	}

	McksNodeConfig struct {
		Connection string `json:"connection"`
		Count      int    `json:"count"`
		Spec       string `json:"spec"`
	}

	McksNodeReq struct {
		ControlPlane []McksNodeConfig `json:"controlPlane"`
		Worker       []McksNodeConfig `json:"worker"`
	}

/*
	Nodes struct {
		Model
		clusterName string
	}
*/
)

/*
func NewNodes(namespace, clusterName string) *Nodes {
	return &Nodes{
		Model:       Model{namespace: namespace},
		clusterName: clusterName,
	}
}
*/
func (self *Mcks) AddNodes(clusterName string, req *McksNodeReq) (*McksNodeList, error) {
	var resp McksNodeList

	_, err := self.execute(
		http.MethodPost,
		fmt.Sprintf("/ns/%s/clusters/%s/nodes", self.namespace, clusterName),
		req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (self *Mcks) GetNode(clusterName, nodeName string) (*McksNode, error) {
	var resp McksNode

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("/ns/%s/clusters/%s/nodes/%s", self.namespace, clusterName, nodeName),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil
}

func (self *Mcks) RemoveNode(clusterName, nodeName string) (*McksStatus, error) {
	var status McksStatus

	cluster, err := self.GetNode(clusterName, nodeName)
	if err != nil {
		return nil, err
	}
	if cluster != nil {
		_, err := self.execute(
			http.MethodDelete,
			fmt.Sprintf("/ns/%s/clusters/%s/nodes/%s", self.namespace, clusterName, nodeName),
			nil, &status)
		if err != nil {
			return nil, err
		}
	} else {
		logger.Infof("Cannot delete the node (namespace=%s, cluster=%s, name=%s, cause=not found)",
			self.namespace, clusterName, nodeName)
	}

	return &status, nil
}
