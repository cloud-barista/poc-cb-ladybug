// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista

package rest_api

import (
	"github.com/cloud-barista/cb-fw-template/pkg/rest-api/common"

	//"os"

	"fmt"

	// REST API (echo)
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// CB-Store
)

/*
// CB-Store
var cblog *logrus.Logger
var store icbs.Store

func init() {
	cblog = config.Cblogger
	store = cbstore.GetStore()
}

type KeyValue struct {
	Key   string
	Value string
}
*/

//var masterConfigInfos confighandler.MASTERCONFIGTYPE

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	Version = " Version: Espresso"
	website = " Repository: https://github.com/cloud-barista/cb-fw-template"
	banner  = `CB-Myfw`
)

// Main Body

func ApiServer() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! This is CB-Myfw")
	})
	e.HideBanner = true
	//e.colorer.Printf(banner, e.colorer.Red("v"+Version), e.colorer.Blue(website))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Printf(banner)
	fmt.Println("")
	fmt.Printf(ErrorColor, Version)
	fmt.Println("")
	fmt.Printf(InfoColor, website)
	fmt.Println("")
	fmt.Println("")

	// Route
	g := e.Group("/myfw", common.ApiValidation())

	g.POST("/myObject", common.RestPostObject)
	g.GET("/myObject", common.RestGetAllObjects)
	g.GET("/myObject/:objectId", common.RestGetObject)
	g.DELETE("/myObject/:objectId", common.RestDeleteObject)

	e.Logger.Fatal(e.Start(":4321"))

}

var SPIDER_URL string
