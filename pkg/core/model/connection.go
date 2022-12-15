package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/lang"
)

type Connection struct {
	Model
	namespace string
}

type ConnectionList struct {
	ListModel
	namespace string
	Items     []Connection `json:"items"`
}

func NewConnection(namespace string, name string) *Connection {
	return &Connection{
		Model:     Model{Kind: KIND_CONNECTION, Name: name},
		namespace: namespace,
	}
}

func NewConnectionList(namespace string) *ConnectionList {
	return &ConnectionList{
		ListModel: ListModel{Kind: KIND_CONNECTION_LIST},
		namespace: namespace,
		Items:     []Connection{},
	}
}

func (self *Connection) Select() error {
	key := lang.GetStoreConnectionKey(self.namespace, self.Name)
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

func (self *Connection) Insert() error {
	key := lang.GetStoreConnectionKey(self.namespace, self.Name)
	value, _ := json.Marshal(self)
	err := common.CBStore.Put(key, string(value))
	if err != nil {
		return err
	}

	return nil
}

func (self *Connection) Delete() error {
	key := lang.GetStoreConnectionKey(self.namespace, self.Name)
	err := common.CBStore.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (self *ConnectionList) SelectList() error {
	keyValues, err := common.CBStore.GetList(lang.GetStoreConnectionKey(self.namespace, ""), true)
	if err != nil {
		return err
	}
	for _, keyValue := range keyValues {
		conn := &Connection{namespace: self.namespace}
		json.Unmarshal([]byte(keyValue.Value), &conn)
		self.Items = append(self.Items, *conn)
	}

	return nil
}
