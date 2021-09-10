package chartmuseum

import (
	"fmt"
	"net/http"
	"os"
)

type Chartmuseum struct {
	Model
}

func NewChartmuseum(repoName string) *Chartmuseum {
	return &Chartmuseum{
		Model: Model{repoName: repoName},
	}
}

func (self *Chartmuseum) GetAllCharts() (*map[string][]CmChart, error) {
	var resp map[string][]CmChart

	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return nil, err
	}

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("%s/api/charts", repoUrl),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil
}

func (self *Chartmuseum) GetAllVersionsOfChart(chartName string) (*[]CmChart, error) {
	var resp []CmChart

	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return nil, err
	}

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("%s/api/charts/%s", repoUrl, chartName),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil
}

func (self *Chartmuseum) GetChart(chartName, version string) (*CmChart, error) {
	var resp CmChart

	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return nil, err
	}

	found, err := self.execute(
		http.MethodGet,
		fmt.Sprintf("%s/api/charts/%s/%s", repoUrl, chartName, version),
		nil, &resp)
	if err != nil {
		return nil, err
	}

	if found == false {
		return nil, nil
	}

	return &resp, nil

}

func (self *Chartmuseum) UploadChart(pkgPath string) error {
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
	/*
		resp, err := app.ExecuteHTTP(
			http.MethodPost,
			fmt.Sprintf("%s/api/charts", repoUrl),
			bytesPkg, nil)
		if err != nil {
			return err
		}

		return handlePushResponse(resp.RawResponse)
	*/
	_, err = self.execute(
		http.MethodPost,
		fmt.Sprintf("%s/api/charts", repoUrl),
		bytesPkg, nil)
	if err != nil {
		return err
	}

	return nil
}

func (self *Chartmuseum) DeleteChart(chartName, version string) error {
	repoUrl, err := getRepoUrl(self.repoName)
	if err != nil {
		return err
	}

	if version == "" {
		return fmt.Errorf("The version should be specified when deleting a chart")
	}

	url := fmt.Sprintf("%s/api/charts/%s/%s", repoUrl, chartName, version)

	_, err = self.execute(
		http.MethodDelete,
		url,
		nil, nil)
	if err != nil {
		return err
	}

	return nil
}
