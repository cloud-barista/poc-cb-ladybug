package service

import (
	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	hc "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
)

const (
	defaultHelmCachePath      = "./.go-helm-client/.helmcache"
	defaultHelmRepoConfigPath = "./.go-helm-client/.helmrepo"
)

func getHelmClient(namespace string) (hc.Client, error) {
	// Create global HelmClient for repository
	opt := &hc.Options{
		Namespace:        namespace,
		RepositoryCache:  defaultHelmCachePath,
		RepositoryConfig: defaultHelmRepoConfigPath,
		Debug:            true,
		Linting:          true,
	}

	hcGeneral, err := hc.New(opt)
	if err != nil {
		common.CBLog.Errorf(err.Error())
		return nil, err
	}

	return hcGeneral, nil
}

func getHelmClientFromKubeConf(namespace string, kubeConf *string) (hc.Client, error) {
	opt := &hc.KubeConfClientOptions{
		Options: &hc.Options{
			Namespace:        namespace,
			RepositoryCache:  defaultHelmCachePath,
			RepositoryConfig: defaultHelmRepoConfigPath,
			Debug:            true,
			Linting:          true,
		},
		KubeContext: "",
		KubeConfig:  []byte(*kubeConf),
	}

	hcKube, err := hc.NewClientFromKubeConf(opt)
	if err != nil {
		common.CBLog.Errorf(err.Error())
		return nil, err
	}

	return hcKube, nil
}

func addPackageRepo(namespace, url string) error {
	hClient, err := getHelmClient(namespace)
	if err != nil {
		return err
	}
	//
	// Add local chartmuseum helm repository
	//
	chartRepo := repo.Entry{
		Name: namespace,
		URL:  url,
	}

	// Add a chart-repository to the client
	if err := hClient.AddOrUpdateChartRepo(chartRepo); err != nil {
		return err
	}

	return nil
}
