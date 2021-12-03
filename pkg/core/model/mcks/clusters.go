package mcks

import (
	"fmt"
	"net/http"
)

const (
	MCKS_CLUSTER_STATUS_PHASE_PENDING      = "Pending"
	MCKS_CLUSTER_STATUS_PHASE_PROVISIONING = "Provisioning"
	MCKS_CLUSTER_STATUS_PHASE_PROVISIONED  = "Provisioned"
	MCKS_CLUSTER_STATUS_PHASE_FAILED       = "Failed"
	MCKS_CLUSTER_STATUS_PHASE_DELETING     = "Deleting"
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
		Config          McksConfig       `json:"config"`
		ControlPlane    []McksNodeConfig `json:"controlPlane"`
		Description     string           `json:"description"`
		InstallMonAgent string           `json:"installMonAgent"`
		Label           string           `json:"label"`
		Name            string           `json:"name"`
		Worker          []McksNodeConfig `json:"worker"`
	}

	McksCluster struct {
		ClusterConfig   string     `json:"clusterConfig"`
		CpLeader        string     `json:"cpLeader"`
		CreatedTime     string     `json:"createdTime"`
		Description     string     `json:"description"`
		InstallMonAgent string     `json:"installMonAgent"`
		Kind            string     `json:"kind"`
		Label           string     `json:"label"`
		Mcis            string     `json:"mcis"`
		Name            string     `json:"name"`
		Namespace       string     `json:"namespace"`
		NetworkCni      string     `json:"networkCni"`
		Nodes           []McksNode `json:"nodes"`
		Status          struct {
			Phase   string `json:"phase"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"status"`
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

func (self *Mcks) GetCluster(name string) (*McksCluster, error) {
	var resp McksCluster

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("/ns/%s/clusters/%s", self.namespace, name),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil
}

func (self *Mcks) DeleteCluster(name string) (*McksStatus, error) {
	var status McksStatus

	/*
			cluster, err := self.GetCluster(name)
			if err != nil {
				return nil, err
			}
		if cluster != nil {
	*/
	_, err := self.execute(
		http.MethodDelete,
		fmt.Sprintf("/ns/%s/clusters/%s", self.namespace, name),
		nil, &status)
	if err != nil {
		return nil, err
	}
	/*
		} else {
			common.CBLog.Infof("MCKS: cannot delete the cluster (namespace=%s, name=%s, cause=not found)",
				self.namespace, name)
		}
	*/
	return &status, nil
}
