package chartmuseum

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
)

type ChartsMap struct {
	repoName      string
	RespChartsMap map[string][]CmChart
}

func NewChartsMap(name string) (*ChartsMap, error) {
	if name == "" {
		return nil, errors.New("no valid repository name")
	}

	return &ChartsMap{
		repoName: name,
	}, nil
}

func (self *ChartsMap) GET() (bool, error) {
	url, err := getRepoUrl(self.repoName)
	if err != nil {
		return false, err
	}

	resp, err := app.ExecuteHTTP(
		http.MethodGet,
		fmt.Sprintf("%s/api/charts", url),
		nil, &self.RespChartsMap)
	if err != nil {
		return false, err
	}
	if resp.StatusCode() == http.StatusNotFound {
		return false, nil
	}

	return true, nil
}
