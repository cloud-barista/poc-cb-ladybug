package main

import (
	"sync"

	restapi "github.com/cloud-barista/cb-ladybug/pkg/rest-api"
	"github.com/cloud-barista/cb-ladybug/pkg/utils/config"
)

// @title CB-Ladybug REST API
// @version 0.1.0
// @description CB-Ladybug REST API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1592
// @BasePath /ladybug
func main() {

	config.Setup()

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		restapi.RunServer()
		wg.Done()
	}()

	/*
		go func() {
			grpcserver.RunServer()
			wg.Done()
		}()
	*/
	wg.Wait()

}
