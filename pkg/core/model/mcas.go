package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

const (
	STATUS_MCAS_ENABLED  = "enabled"
	STATUS_MCAS_DISABLED = "disabled"
)

type Mcas struct {
	Model
	Namespace string
	Status    string
}

func NewMcas(namespace string) *Mcas {
	return &Mcas{
		Model:     Model{Kind: KIND_MCAS},
		Namespace: namespace,
	}
}

func (self *Mcas) Init() error {
	return self.Enable()
}

func (self *Mcas) GetStatus() (string, error) {
	key := lang.GetStoreMcasKey(self.Namespace)
	keyValue, err := common.CBStore.Get(key)
	if err != nil {
		return "", err
	}
	if keyValue == nil {
		return "", errors.New(fmt.Sprintf("%s not found", key))
	}

	json.Unmarshal([]byte(keyValue.Value), &self)
	return self.Status, nil
}

func (self *Mcas) Enable() error {
	key := lang.GetStoreMcasKey(self.Namespace)
	self.Status = STATUS_MCAS_ENABLED
	value, _ := json.Marshal(self)
	err := common.CBStore.Put(key, string(value))
	if err != nil {
		return err
	}

	return nil
}

func (self *Mcas) Disable() error {
	key := lang.GetStoreMcasKey(self.Namespace)
	self.Status = STATUS_MCAS_DISABLED
	value, _ := json.Marshal(self)
	err := common.CBStore.Put(key, string(value))
	if err != nil {
		return err
	}

	return nil
}

func (self *Mcas) Delete() error {
	key := lang.GetStoreMcasKey(self.Namespace)
	err := common.CBStore.Delete(key)
	if err != nil {
		return err
	}

	return nil
}