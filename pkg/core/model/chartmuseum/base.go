package chartmuseum

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/chartmuseum/helm-push/pkg/helm"
	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
)

// Chartmuseum's Chart
type CmChart struct {
	Name        string         `json:"name"`
	Home        string         `json:"home"`
	Sources     []string       `json:"sources"`
	Version     string         `json:"version"`
	Description string         `json:"description"`
	Maintainers []CmMaintainer `json:"maintainers"`
	Engine      string         `json:"engine"`
	Icon        string         `json:"icon"`
	Urls        []string       `json:"urls"`
	Created     string         `json:"created"`
	Digest      string         `json:"digest"`
}

// CmMaintainer reprecents a chartmuseum's chart maintainer
type CmMaintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CmError struct {
	Error string `json:"error"`
}

type Model struct {
	repoName string
}

func (self *Model) execute(method string, url string, body interface{}, result interface{}) (bool, error) {
	/*
		// validation
		if err := self.validate(validation.Validation{}); err != nil {
			return false, err
		}
	*/
	resp, err := app.ExecuteHTTP(method, url, body, result)
	if err != nil {
		return false, err
	}

	// response check
	if method == http.MethodPost && resp.StatusCode() != http.StatusCreated {
		er := CmError{}
		err := json.Unmarshal(resp.Body(), &er)
		if err != nil || er.Error == "" {
			common.CBLog.Warnf("Chartmuseum: statusCode=%d, url=%s, body=%s", resp.StatusCode(), resp.Request.URL, resp)
			return false, errors.New("Unknown error")
		}
		return false, errors.New(er.Error)
	}

	if resp.StatusCode() > 300 && resp.StatusCode() != http.StatusNotFound {
		common.CBLog.Warnf("Chartmuseum: statusCode=%d, url=%s, body=%s", resp.StatusCode(), resp.Request.URL, resp)
		return false, errors.New("Unknown error")
	}

	if method == http.MethodGet && resp.StatusCode() == http.StatusNotFound {
		common.CBLog.Infof("Not found data (status=404, method=%s, url=%s)", method, url)
		return false, nil
	}

	return true, nil
}

/*
func (self *Model) validate(valid validation.Validation) error {
	valid.Required(self.repoName, "repoName")
	//valid.Required(self.Name, "name")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errors.New(fmt.Sprintf("[%s]%s", err.Key, err.Error()))
		}
	}
	return nil
}
*/

func getRepoUrl(repoName string) (string, error) {
	if repoName == "" {
		return "", errors.New("no valid repository name")
	}
	repo, err := helm.GetRepoByName(repoName)
	if err != nil {
		return "", err
	}

	url := strings.Replace(repo.Config.URL, "cm://", "https://", 1)
	return url, err
}
