package chartmuseum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
	"github.com/go-resty/resty/v2"
)

type Charts struct {
	repoName      string
	name          string
	RespChart     CmChart
	RespChartList []CmChart
}

func NewCharts(repoName, name string) *Charts {
	return &Charts{
		repoName: repoName,
		name:     name,
	}
}

func (self *Charts) GET(version string) (bool, error) {
	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return false, err
	}

	var resp *resty.Response
	if version == "" {
		resp, err = app.ExecuteHTTP(
			http.MethodGet,
			fmt.Sprintf("%s/api/charts/%s", repoUrl, self.name),
			nil, &self.RespChartList)
	} else {
		resp, err = app.ExecuteHTTP(
			http.MethodGet,
			fmt.Sprintf("%s/api/charts/%s/%s", repoUrl, self.name, version),
			nil, &self.RespChart)
	}

	if err != nil {
		return false, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return false, nil
	}

	return true, nil
}

func (self *Charts) POST(pkgPath string) error {
	var err error

	bytesPkg, err := os.ReadFile(pkgPath)
	if err != nil {
		return err
	}

	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return err
	}
	/*
		conf := config.Config

		req := resty.New().SetDisableWarn(true).R().SetBasicAuth(*conf.Username, *conf.Password)
		req.SetBody(bytesPkg)
		req.SetContentLength(true)

		// execute
		resp, err := req.Execute(
			http.MethodPost,
			fmt.Sprintf("%s/api/charts", repoUrl))
		if err != nil {
			return err
		}
	*/

	resp, err := app.ExecuteHTTP(
		http.MethodPost,
		fmt.Sprintf("%s/api/charts", repoUrl),
		bytesPkg, nil)
	if err != nil {
		return err
	}

	return handlePushResponse(resp.RawResponse)
}

/*
func (self *Charts) POST(pkgPath string) error {
	var err error

	f, err := os.Stat(pkgPath)
	if err != nil {
		return err
	}

	if !f.Mode().IsRegular() {
		return errors.New("no valid package")
	}

	url, err := getRepoUrl(self.repoName)
	if err != nil {
		return err
	}

	client, err := cm.NewClient(
		cm.URL(url),
	)
	if err != nil {
		return err
	}

	resp, err := client.UploadChartPackage(pkgPath, true)
	return handlePushResponse(resp)
}
*/
func handlePushResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusCreated {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return getChartmuseumError(b, resp.StatusCode)
	}
	return nil
}

func getChartmuseumError(b []byte, code int) error {
	var er struct {
		Error string `json:"error"`
	}
	err := json.Unmarshal(b, &er)
	if err != nil || er.Error == "" {
		return fmt.Errorf("%d: could not properly parse response JSON: %s", code, string(b))
	}
	return fmt.Errorf("%d: %s", code, er.Error)
}

func (self *Charts) DELETE(version string) error {
	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return err
	}

	var url string
	if version == "" {
		return fmt.Errorf("When deleting a chart, the version should be specified")
	} else {
		url = fmt.Sprintf("%s/api/charts/%s/%s", repoUrl, self.name, version)
	}

	_, err = app.ExecuteHTTP(
		http.MethodDelete,
		url,
		nil, nil)
	if err != nil {
		return err
	}

	return nil
}
