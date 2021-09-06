package chartmuseum

import (
	"errors"
	"strings"

	"github.com/chartmuseum/helm-push/pkg/helm"
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
