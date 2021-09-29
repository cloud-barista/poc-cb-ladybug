package chartmuseum

import (
	"fmt"
	"net/http"
)

type Chartmuseum struct {
	Model
}

func NewChartmuseum(repoName string) *Chartmuseum {
	return &Chartmuseum{
		Model: Model{RepoName: repoName},
	}
}

/*
func (self *Chartmuseum) GetRepo() string {
	return self.RepoName
}
*/

func (self *Chartmuseum) GetAllCharts() (*map[string][]CmChart, error) {
	var resp map[string][]CmChart

	repoUrl, err := getRepoUrl(self.RepoName)
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

func (self *Chartmuseum) GetChartAllVersions(chartName string) (*[]CmChart, error) {
	var resp []CmChart

	repoUrl, err := getRepoUrl(self.RepoName)
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

	repoUrl, err := getRepoUrl(self.RepoName)
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

func (self *Chartmuseum) UploadChart(bytesPkg []byte) error {
	repoUrl, err := getRepoUrl(self.RepoName)
	if err != nil {
		return err
	}

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
	repoUrl, err := getRepoUrl(self.RepoName)
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
