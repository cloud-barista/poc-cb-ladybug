package mcks

import (
	"fmt"
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

type (
	McksKubernetes struct {
		NetworkCni       string `json:"networkCni"`
		PodCidr          string `json:"podCidr"`
		ServiceCidr      string `json:"serviceCidr"`
		ServiceDnsDomain string `json:"serviceDnsDomain"`
	}

	McksConfig struct {
		Kubernetes McksKubernetes `json:"kubernetes"`
	}

	McksClusterReq struct {
		Config       McksConfig       `json:"config"`
		ControlPlane []McksNodeConfig `json:"controlPlane"`
		Name         string           `json:"name"`
		Worker       []McksNodeConfig `json:"worker"`
	}

	McksCluster struct {
		ClusterConfig string     `json:"clusterConfig"`
		CpLeader      string     `json:"cpLeader"`
		Kind          string     `json:"kind"`
		Mcis          string     `json:"mcis"`
		Name          string     `json:"name"`
		Namespace     string     `json:"namespace"`
		NetworkCni    string     `json:"networkCni"`
		Nodes         []McksNode `json:"nodes"`
		Status        string     `json:"status"`
		Uid           string     `json:"uid"`
	}

	McksClusterList struct {
		Kind  string        `json:"kind"`
		Items []McksCluster `json:"items"`
	}
)

func (self *Mcks) ListCluster() (*McksClusterList, error) {
	var resp McksClusterList

	_, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("/ns/%s/clusters", self.namespace),
		nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (self *Mcks) CreateCluster(req McksClusterReq) (*McksCluster, error) {
	var resp McksCluster

	_, err := self.execute(
		http.MethodPost,
		fmt.Sprintf("/ns/%s/clusters", self.namespace),
		req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (self *Mcks) GetCluster(clusterName string) (*McksCluster, error) {
	var resp McksCluster

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("/ns/%s/clusters/%s", self.namespace, clusterName),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil
}

func (self *Mcks) DeleteCluster(clusterName string) (*McksStatus, error) {
	var status McksStatus

	cluster, err := self.GetCluster(clusterName)
	if err != nil {
		return nil, err
	}
	if cluster != nil {
		_, err := self.execute(
			http.MethodDelete,
			fmt.Sprintf("/ns/%s/clusters/%s", self.namespace, clusterName),
			nil, &status)
		if err != nil {
			return nil, err
		}
	} else {
		logrus.Infof("Cannot delete the cluster (namespace=%s, name=%s, cause=not found)",
			self.namespace, clusterName)
	}

	return &status, nil
}
