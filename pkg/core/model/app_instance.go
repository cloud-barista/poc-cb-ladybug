package model

import (
	"errors"
	"time"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

type (
	AppInstance struct {
		//	Model
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		PackageName string `json:"packageName"`
		Version     string `json:"version"`
	}

	AppInstanceList struct {
		//ListModel
		Namespace string
		Items     []AppInstance `json:"items"`
	}

	AppInstanceReq struct {
		InstanceName string        `json:"instance"`
		PackageName  string        `json:"package"`
		Version      string        `json:"version,omitempty"`
		Wait         bool          `json:"wait,omitempty"`
		Timeout      time.Duration `json:timeout,omitempty"`
		Force        bool          `json:"force,omitempty"`
		UpgradeCRDs  bool          `json:"upgradeCRDs,omitempty"`
	}
)

func NewAppInstance(namespace, instName, pkgName, version string) *AppInstance {
	return &AppInstance{
		Name:        instName,
		Namespace:   namespace,
		PackageName: pkgName,
		Version:     version,
	}
}

func NewAppInstanceList(namespace string) *AppInstanceList {
	return &AppInstanceList{
		Namespace: namespace,
		Items:     []AppInstance{},
	}
}

/*
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
*/
func AppInstanceReqDef(req *AppInstanceReq) {
	printAppInstanceReq(req)
	if req.Timeout == 0 {
		// Set timeout to helm default timeout(300 second)
		req.Timeout = 300
	}
	printAppInstanceReq(req)
}

func AppInstanceReqValidate(req *AppInstanceReq) error {
	if len(req.InstanceName) == 0 {
		return errors.New("app instance name is empty")
	} else {
		err := lang.CheckName(req.InstanceName)
		if err != nil {
			return err
		}
	}

	if len(req.PackageName) == 0 {
		return errors.New("app package name is empty")
	} else {
		err := lang.CheckName(req.PackageName)
		if err != nil {
			return err
		}
	}

	return nil
}

func printAppInstanceReq(req *AppInstanceReq) {
	common.CBLog.Debugf("AppInstanceReq.InstanceName:\t%v", req.InstanceName)
	common.CBLog.Debugf("AppInstanceReq.PackageName:\t%v", req.PackageName)
	common.CBLog.Debugf("AppInstanceReq.Version:\t%v", req.Version)
	common.CBLog.Debugf("AppInstanceReq.Wait:\t\t%v", req.Wait)
	common.CBLog.Debugf("AppInstanceReq.Timeout:\t%v", req.Timeout*time.Second)
	common.CBLog.Debugf("AppInstanceReq.Force:\t\t%v", req.Force)
	common.CBLog.Debugf("AppInstanceReq.UpgradeCRDs:\t%v", req.UpgradeCRDs)
}
