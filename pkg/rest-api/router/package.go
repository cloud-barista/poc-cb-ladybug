package router

import (
	"net/http"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/service"
	"github.com/labstack/echo/v4"
)

func ListPackage(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	pkgList, err := service.ListPackage(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, pkgList)
}

func GetPackageAllVersions(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "package"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	pkg, err := service.GetPackageAllVersions(c.Param("namespace"), c.Param("package"))
	if err != nil {
		common.CBLog.Infof("not found a package (namespace=%s, pakcage=%s, cause=%s)", c.Param("namespace"), c.Param("package"), err)
		return SendMessage(c, http.StatusNotFound, err.Error())
	}

	return Send(c, http.StatusOK, pkg)
}

func GetPackage(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "package", "version"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	pkg, err := service.GetPackage(c.Param("namespace"), c.Param("package"), c.Param("version"))
	if err != nil {
		common.CBLog.Infof("not found a package (namespace=%s, pakcage=%s, version=%s, cause=%s)", c.Param("namespace"), c.Param("package"), c.Param("version"), err)
		return SendMessage(c, http.StatusNotFound, err.Error())
	}

	return Send(c, http.StatusOK, pkg)
}

func UploadPackage(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	namespace := c.Param("namespace")
	fhPkg, err := c.FormFile("package")
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}
	status, err := service.UploadPackage(namespace, fhPkg)
	if err != nil {
		common.CBLog.Error(err)
		return Send(c, status.Code, status)
	}

	return Send(c, http.StatusOK, status)
}

func DeletePackage(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "package", "version"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	status, err := service.DeletePackage(c.Param("namespace"), c.Param("package"), c.Param("version"))
	if err != nil {
		common.CBLog.Infof("not found a package (namespace=%s, pakcage=%s, version=%s, cause=%s)", c.Param("namespace"), c.Param("package"), c.Param("version"), err)
		return SendMessage(c, http.StatusNotFound, err.Error())
	}

	return Send(c, http.StatusOK, status)
}
