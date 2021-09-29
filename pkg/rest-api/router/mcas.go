package router

import (
	"net/http"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/core/model"
	_ "github.com/cloud-barista/cb-mcas/pkg/core/service"
	"github.com/cloud-barista/cb-mcas/pkg/utils/app"
	"github.com/labstack/echo/v4"
)

func GetMcas(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	mcas := model.NewMcas(c.Param("namespace"))
	status, err := mcas.GetStatus()
	if err != nil {
		common.CBLog.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return app.Send(c, http.StatusOK, status)
}

func SetMcas(c echo.Context) error {
	/*
		mcas := model.NewMcas(c.Param("namespace"))
		status, err := mcas.GetStatus()
		if err != nil {
			common.CBLog.Error(err)
			return app.SendMessage(c, http.StatusBadRequest, err.Error())
		}

		return app.Send(c, http.StatusOK, status)
	*/
	return nil
}
