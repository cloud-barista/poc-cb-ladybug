// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista

package restapi

import (
	"crypto/subtle"
	"fmt"

	"github.com/cloud-barista/cb-mcas/pkg/core/common"
	"github.com/cloud-barista/cb-mcas/pkg/rest-api/router"
	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	cfg "github.com/cloud-barista/cb-mcas/pkg/utils/config"

	// REST API (echo)
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// CB-Store

	//_ "github.com/cloud-barista/cb-mcas/pkg/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//var masterConfigInfos confighandler.MASTERCONFIGTYPE

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	Version = " Version: Affogato"
	website = " Repository: https://github.com/cloud-barista/poc-cb-ladybug"
	banner  = `CB-Ladybug`
)

// Main Body
func RunServer() {
	// Echo instance
	e := echo.New()

	// Echo middleware func
	e.Use(middleware.Logger())  // Setting logger
	e.Use(middleware.Recover()) // Recover from panics anywhere in the chain
	/*
		if *config.Config.LoglevelHTTP == true {
		}
	*/

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ // CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(*cfg.Config.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(*cfg.Config.Password)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	fmt.Println("\n \n ")
	fmt.Printf(banner)
	fmt.Println("\n ")
	fmt.Printf(ErrorColor, Version)
	fmt.Println("")
	fmt.Printf(InfoColor, website)
	fmt.Println("\n \n ")

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET(*config.Config.BasePath+"/health", router.Health)

	g := e.Group(*config.Config.BasePath+"/ns", common.NsValidate())

	// Routes
	g.POST("/:namespace/mcas", router.EnableMcas)
	g.GET("/:namespace/mcas", router.GetMcas)
	g.DELETE("/:namespace/mcas", router.DisableMcas)

	g.POST("/:namespace/packages", router.UploadPackage)
	g.GET("/:namespace/packages", router.ListPackage)
	g.GET("/:namespace/packages/:package", router.GetPackageAllVersions)
	g.GET("/:namespace/packages/:package/:version", router.GetPackage)
	g.DELETE("/:namespace/packages/:package/:version", router.DeletePackage)

	g.POST("/:namespace/clusters", router.CreateCluster)
	g.GET("/:namespace/clusters", router.ListCluster)
	g.GET("/:namespace/clusters/:cluster", router.GetCluster)
	g.DELETE("/:namespace/clusters/:cluster", router.DeleteCluster)

	// Start server
	e.Logger.Fatal(e.Start(*cfg.Config.ListenAddress))
}
