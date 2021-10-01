package service

import (
	"strings"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
)

func GetMcas(namespace string) (string, error) {
	mcas := model.NewMcas(namespace)
	status, err := mcas.GetStatus()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err2 := mcas.Disable()
			if err2 != nil {
				common.CBLog.Error(err2)
				return "", err2
			}
			status, err2 = mcas.GetStatus()
			if err2 != nil {
				common.CBLog.Error(err2)
				return "", err2
			}
		} else {
			common.CBLog.Error(err)
			return "", err
		}
	}

	return string(status), nil
}

func EnableMcas(namespace string) error {
	mcas := model.NewMcas(namespace)
	err := mcas.Enable()
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	return nil
}

func DisableMcas(namespace string) error {
	mcas := model.NewMcas(namespace)
	err := mcas.Disable()
	if err != nil {
		common.CBLog.Error(err)
		return err
	}

	return nil
}
