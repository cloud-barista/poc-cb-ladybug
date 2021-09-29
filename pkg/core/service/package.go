package service

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
	"github.com/cloud-barista/cb-mcas/pkg/core/model/chartmuseum"
)

func ListPackage(namespace string) (*model.PackageList, error) {
	pkglist := model.NewPackageList(namespace)

	cm := chartmuseum.NewChartmuseum(namespace)

	allcharts, err := cm.GetAllCharts()
	if err != nil {
		return nil, err
	}

	for _, charts := range *allcharts {
		for _, ch := range charts {
			pkg := model.NewPackage(cm.RepoName, ch.Name, ch.Version)
			pkglist.Items = append(pkglist.Items, *pkg)
		}
	}

	return pkglist, nil
}

func GetPackageAllVersions(namespace, name string) (*[]model.Package, error) {
	common.CBLog.Debugf("[CALLED]")

	var pkgs []model.Package

	cm := chartmuseum.NewChartmuseum(namespace)
	charts, err := cm.GetChartAllVersions(name)
	if err != nil {
		return nil, err
	}

	for _, ch := range *charts {
		pkg := model.NewPackage(cm.RepoName, ch.Name, ch.Version)
		pkgs = append(pkgs, *pkg)
	}

	return &pkgs, nil
}

func GetPackage(namespace, name, version string) (*model.Package, error) {
	common.CBLog.Debugf("[CALLED]")

	cm := chartmuseum.NewChartmuseum(namespace)
	chart, err := cm.GetChart(name, version)
	if err != nil {
		return nil, err
	}

	pkg := model.NewPackage(cm.RepoName, chart.Name, chart.Version)
	return pkg, nil
}

func UploadPackage(namespace string, fhPkg *multipart.FileHeader) (*model.Status, error) {
	common.CBLog.Debugf("[CALLED]")

	status := model.NewStatus()
	status.Code = model.STATUS_FAIL

	//	name := strings.TrimSuffix(fhPkg.Filename, path.Ext(fhPkg.Filename))

	filePkg, err := fhPkg.Open()
	if err != nil {
		common.CBLog.Error(err)
		return status, err
	}
	defer filePkg.Close()

	cm := chartmuseum.NewChartmuseum(namespace)

	bytesPkg, err := ioutil.ReadAll(filePkg)
	if err != nil {
		common.CBLog.Error(err)
		return status, err
	}

	err = cm.UploadChart(bytesPkg)
	if err != nil {
		common.CBLog.Error(err)
		if err.Error() == "file already exists" {
			status.Message = "package already exists"
		}
		return status, err
	}

	status.Code = model.STATUS_SUCCESS
	status.Message = "package is uploaded"
	return status, nil
}

func DeletePackage(namespace, name, version string) (*model.Status, error) {
	status := model.NewStatus()
	status.Code = model.STATUS_FAIL

	cm := chartmuseum.NewChartmuseum(namespace)
	err := cm.DeleteChart(name, version)
	if err != nil {
		common.CBLog.Error(err)
		return nil, err
	}

	status.Code = model.STATUS_SUCCESS
	return status, nil
}
