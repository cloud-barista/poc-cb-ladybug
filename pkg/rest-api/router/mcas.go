package router

import (
	"net/http"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/service"
	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
	"github.com/labstack/echo/v4"
)

func GetMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	status, err := service.GetMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return app.Send(c, http.StatusOK, string(status))
}

func EnableMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	err := service.EnableMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return app.Send(c, http.StatusOK, nil)
}

func DisableMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	err := service.DisableMcas(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return app.Send(c, http.StatusOK, nil)
}
