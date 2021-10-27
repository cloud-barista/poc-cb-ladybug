package router

import (
	"net/http"

	"github.com/cloud-barista/cb-ladybug/pkg/core/common"
	"github.com/labstack/echo/v4"
)

// Health Method
// @Tags Default
// @Summary Health Check
// @Description for health check
// @ID Health
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /health [get]
func Health(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")
	return c.String(http.StatusOK, "cloud-barista cb-ladybug is alived\n")
}
