package router

import (
	"net/http"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/cb-ladybug/pkg/core/service"
	"github.com/labstack/echo/v4"
)

func GetMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	status, err := service.GetMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, status)
}

func EnableMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	err := service.EnableMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, "enabled")
}

func DisableMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	err := service.DisableMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, "disabled")
}
