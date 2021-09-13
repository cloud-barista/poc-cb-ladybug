package tumblebug

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
	"github.com/cloud-barista/cb-mcas/pkg/utils/config"

	logger "github.com/sirupsen/logrus"
)

type (
	TumblebugStatus struct {
		Message string `json:"message"`
	}
	Tumblebug struct {
		Model
	}
)

func NewTumblebug() *Tumblebug {
	return &Tumblebug{}
}

type Model struct {
}

func (self *Model) execute(method string, url string, body interface{}, result interface{}) (bool, error) {

	// validation
	if err := self.validate(validation.Validation{}); err != nil {
		return false, err
	}

	resp, err := app.ExecuteHTTP(method, *config.Config.TumblebugUrl+url, body, result)
	if err != nil {
		return false, err
	}

	// response check
	if resp.StatusCode() > 300 && resp.StatusCode() != http.StatusNotFound {
		logger.Warnf("MCKS: statusCode=%d, url=%s, body=%s", resp.StatusCode(), resp.Request.URL, resp)
		status := TumblebugStatus{}
		json.Unmarshal(resp.Body(), &status)
		/*
			// message > message 로 리턴되는 경우가 있어서 한번더 unmarshal 작업
			if json.Valid([]byte(status.Message)) {
				json.Unmarshal([]byte(status.Message), &status)
			}
		*/
		return false, errors.New(status.Message)
	}

	if method == http.MethodGet && resp.StatusCode() == http.StatusNotFound {
		logger.Infof("Not found data (status=404, method=%s, url=%s)", method, url)
		return false, nil
	}

	return true, nil
}

func (self *Model) validate(valid validation.Validation) error {
	//valid.Required(self.namespace, "namespace")
	//valid.Required(self.Name, "name")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errors.New(fmt.Sprintf("[%s]%s", err.Key, err.Error()))
		}
	}
	return nil
}
