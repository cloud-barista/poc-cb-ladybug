package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/cb-ladybug/pkg/utils/lang"
)

type ClusterState int

const (
	CS_CREATING ClusterState = iota + 1
	CS_RUNNING
	CS_FAILED
	CS_UNKNOWN
	CS_DELETED
)

func (s ClusterState) String() string {
	return [...]string{"creating", "running", "failed", "unknown", "deleted"}[s]
}

type (
	ClusterInfo struct {
		Namespace     string       `json:"namespace"`
		Name          string       `json:"name"`
		State         ClusterState `json:"state"`
		ClusterConfig string       `json:"clusterconfig"`
	}

	ClusterInfoList struct {
		Namespace string
		Items     []ClusterInfo
	}
)

func NewClusterInfo(namespace, name string) *ClusterInfo {
	return &ClusterInfo{
		Namespace: namespace,
		Name:      name,
		State:     CS_UNKNOWN,
	}
}

func NewClusterInfoList(namespace string) *ClusterInfoList {
	return &ClusterInfoList{
		Namespace: namespace,
		Items:     []ClusterInfo{},
	}
}

func (self *ClusterInfo) setState(clusterState ClusterState) {
	self.State = clusterState
}

func (self *ClusterInfo) SetClusterConfig(clusterConfig string) {
	self.ClusterConfig = clusterConfig
}

func (self *ClusterInfo) TriggerCreate() error {
	self.setState(CS_CREATING)
	return self.saveToStore()
}

func (self *ClusterInfo) TriggerSuccess() error {
	self.setState(CS_RUNNING)
	return self.saveToStore()
}

func (self *ClusterInfo) TriggerFail() error {
	self.setState(CS_FAILED)
	return self.saveToStore()
}

func (self *ClusterInfo) TriggerUnknown() error {
	self.setState(CS_UNKNOWN)
	return self.saveToStore()
}

func (self *ClusterInfo) TriggerDelete() error {
	self.setState(CS_DELETED)
	return self.saveToStore()
}

func (self *ClusterInfo) saveToStore() error {
	key := lang.GetStoreClusterKey(self.Namespace, self.Name)
	value, _ := json.Marshal(self)
	err := common.CBStore.Put(key, string(value))
	if err != nil {
		return err
	}
	return nil
}

func (self *ClusterInfo) Select() error {
	key := lang.GetStoreClusterKey(self.Namespace, self.Name)
	keyValue, err := common.CBStore.Get(key)
	if err != nil {
		return err
	}
	if keyValue == nil {
		return errors.New(fmt.Sprintf("%s not found", key))
	}
	json.Unmarshal([]byte(keyValue.Value), &self)

	return nil
}

func (self *ClusterInfo) Delete() error {
	// delete cluster
	key := lang.GetStoreClusterKey(self.Namespace, self.Name)
	err := common.CBStore.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (self *ClusterInfoList) SelectList() error {
	keyValues, err := common.CBStore.GetList(lang.GetStoreClusterKey(self.Namespace, ""), true)
	if err != nil {
		return err
	}
	self.Items = []ClusterInfo{}
	for _, kv := range keyValues {
		clusterInfo := &ClusterInfo{}
		json.Unmarshal([]byte(kv.Value), &clusterInfo)

		self.Items = append(self.Items, *clusterInfo)
	}

	return nil

}
