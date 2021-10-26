package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
	hc "github.com/mittwald/go-helm-client"
)

func ListAppInstance(namespace string) (*model.AppInstanceList, error) {
	appInstList := model.NewAppInstanceList(namespace)

	clusterInfoList := model.NewClusterInfoList(namespace)
	clusterInfoList.SelectList()

	for _, cInfo := range clusterInfoList.Items {
		if cInfo.State == model.CS_RUNNING {
			appList, err := listAppInstance(namespace, &cInfo.ClusterConfig)
			if err != nil {
				return nil, err
			}

			appInstList.Items = append(appInstList.Items, appList.Items...)
		}
	}

	return appInstList, nil
}

func GetAppInstance(namespace, name string) (*model.AppInstance, error) {
	clusterInfoList := model.NewClusterInfoList(namespace)
	clusterInfoList.SelectList()

	var selectedCInfo *model.ClusterInfo = nil
	for _, cInfo := range clusterInfoList.Items {
		if cInfo.State == model.CS_RUNNING {
			selectedCInfo = &cInfo
			break
		}
	}

	if selectedCInfo == nil {
		return nil, errors.New(fmt.Sprintf("no available cluster to get the application '%s'", name))
	}

	appInstance, err := getAppInstance(namespace, &selectedCInfo.ClusterConfig, name)
	if err != nil {
		return nil, err
	}

	return appInstance, nil
}

func CreateAppInstance(namespace string, req *model.AppInstanceReq) (*model.AppInstance, error) {
	//
	// TODO: Find an available cluster in namespace
	//
	/*
		  mcks := m.NewMcks(namespace)
			mcksClusterList, err := mcks.ListCluster()
			if err != nil {
				return nil, err
			}

			for _, mcksCluster := range mcksClusterList.Items {
			}
	*/

	clusterInfoList := model.NewClusterInfoList(namespace)
	clusterInfoList.SelectList()

	var selectedCInfo *model.ClusterInfo = nil
	for _, cInfo := range clusterInfoList.Items {
		if cInfo.State == model.CS_RUNNING {
			selectedCInfo = &cInfo
			break
		}
	}

	if selectedCInfo != nil {
		//
		// Install the application to the selected cluster
		//
		err := installAppInstance(namespace, &selectedCInfo.ClusterConfig, req)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("no available cluster to deploy the app '%s'", req.InstanceName))
	}

	return nil, nil
}

func DeleteAppInstance(namespace, name string) (*model.Status, error) {
	status := model.NewStatus()

	clusterInfoList := model.NewClusterInfoList(namespace)
	clusterInfoList.SelectList()

	var selectedCInfo *model.ClusterInfo = nil
	for _, cInfo := range clusterInfoList.Items {
		if cInfo.State == model.CS_RUNNING {
			selectedCInfo = &cInfo
			break
		}
	}

	if selectedCInfo != nil {
		//
		// Uninstall the application from the selected cluster
		//

		common.CBLog.Infof("try to uninstall the application '%s'", name)

		err := uninstallAppInstance(namespace, &selectedCInfo.ClusterConfig, name)
		if err != nil {
			status.Code = model.STATUS_FAIL
			status.Message = fmt.Sprintf("cannot uninstall the application '%s' with error='%s'", name, err.Error())
		} else {
			status.Code = model.STATUS_SUCCESS
			status.Message = fmt.Sprintf("the application '%s' has been uninstalled", name)
		}
	} else {
		common.CBLog.Infof("cannot uninstall the application '%s' because we can not find the cluster", name)

		status.Code = model.STATUS_FAIL
		status.Message = fmt.Sprintf("cannot find the cluster which installed the application '%s'", name)
	}

	return status, nil
}

func listAppInstance(namespace string, kubeConf *string) (*model.AppInstanceList, error) {
	hcKube, err := getHelmClientFromKubeConf(namespace, kubeConf)
	if err != nil {
		return nil, err
	}

	rels, err := hcKube.ListDeployedReleases()
	if err != nil {
		return nil, err
	}

	appInstList := model.NewAppInstanceList(namespace)
	for _, rel := range rels {
		common.CBLog.Debugf("Instance Name: %s, Package Name: %s, Version: %v",
			rel.Name, rel.Chart.Metadata.Name, rel.Chart.Metadata.Version)

		appInst := model.NewAppInstance(namespace, rel.Name, rel.Chart.Metadata.Name, rel.Chart.Metadata.Version)

		appInstList.Items = append(appInstList.Items, *appInst)
	}

	return appInstList, nil
}

func getAppInstance(namespace string, kubeConf *string, appInstName string) (*model.AppInstance, error) {
	hcKube, err := getHelmClientFromKubeConf(namespace, kubeConf)
	if err != nil {
		return nil, err
	}

	rel, err := hcKube.GetRelease(appInstName)
	if err != nil {
		return nil, err
	}

	appInst := model.NewAppInstance(namespace, rel.Name, rel.Chart.Metadata.Name, rel.Chart.Metadata.Version)

	return appInst, nil
}

func installAppInstance(namespace string, kubeConf *string, req *model.AppInstanceReq) error {
	hcKube, err := getHelmClientFromKubeConf(namespace, kubeConf)
	if err != nil {
		return err
	}

	chartSpec := hc.ChartSpec{
		ReleaseName:     req.InstanceName,
		ChartName:       namespace + "/" + req.PackageName,
		Namespace:       namespace,
		Version:         req.Version,
		CreateNamespace: true,
		Wait:            req.Wait,
		Timeout:         req.Timeout * time.Second,
		UpgradeCRDs:     req.UpgradeCRDs,
		Force:           req.Force,
	}

	common.CBLog.Debugf("ChartSpec.ReleaseName: %s, ChartSpec.ChartName: %s", chartSpec.ReleaseName, chartSpec.ChartName)

	common.CBLog.Infof("try to install the application '%s'", req.InstanceName)
	if _, err := hcKube.InstallOrUpgradeChart(context.Background(), &chartSpec); err != nil {
		return err
	}

	return nil
}

func uninstallAppInstance(namespace string, kubeConf *string, appInstName string) error {
	hcKube, err := getHelmClientFromKubeConf(namespace, kubeConf)
	if err != nil {
		return err
	}

	chartSpec := hc.ChartSpec{
		ReleaseName: appInstName,
		Timeout:     600 * time.Second,
	}

	if err = hcKube.UninstallRelease(&chartSpec); err != nil {
		return err
	}

	return nil
}
