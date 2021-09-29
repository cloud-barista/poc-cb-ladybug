package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

type AppInstance struct {
	Model
	Namespace string `json:"namespace"`
	Version   string `json:"version"`
}

type AppInstanceList struct {
	ListModel
	Namespace string
	Items     []AppInstance `json:"items"`
}

func NewAppInstance(namespace, name, version string) *AppInstance {
	return &AppInstance{
		Model:     Model{Kind: KIND_APP_INSTANCE, Name: name},
		Namespace: namespace,
		Version:   version,
	}
}

func NewAppInstanceList(namespace string) *AppInstanceList {
	return &AppInstanceList{
		ListModel: ListModel{Kind: KIND_APP_INSTANCE_LIST},
		Namespace: namespace,
		Items:     []AppInstance{},
	}
}

func (self *AppInstance) Select() error {
	key := lang.GetStoreAppInstanceKey(self.Namespace, self.Name)
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

func (self *AppInstance) Insert() error {
	key := lang.GetStoreAppInstanceKey(self.Namespace, self.Name)
	value, _ := json.Marshal(self)
	err := common.CBStore.Put(key, string(value))
	if err != nil {
		return err
	}

	return nil
}

func (self *AppInstance) Delete() error {
	key := lang.GetStoreAppInstanceKey(self.Namespace, self.Name)
	err := common.CBStore.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (self *AppInstanceList) SelectList() error {
	keyValues, err := common.CBStore.GetList(lang.GetStoreAppInstanceKey(self.Namespace, ""), true)
	if err != nil {
		return err
	}
	for _, keyValue := range keyValues {
		pkg := &AppInstance{}
		json.Unmarshal([]byte(keyValue.Value), &pkg)
		self.Items = append(self.Items, *pkg)
	}

	return nil
}
