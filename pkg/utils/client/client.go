package client

import (
	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	"github.com/go-resty/resty/v2"
)

func ExecuteHTTP(method string, url string, body interface{}, result interface{}) (*resty.Response, error) {

	conf := config.Config

	client := resty.New()

	client.SetDebug(false)
	req := client.SetDisableWarn(true).R().SetBasicAuth(*conf.Username, *conf.Password)

	if body != nil {
		req.SetBody(body)
	}
	if result != nil {
		req.SetResult(result)
	}

	// execute
	return req.Execute(method, url)
}
