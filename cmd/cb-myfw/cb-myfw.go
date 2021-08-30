package main

import (
	"os"

	rest_api "github.com/cloud-barista/cb-fw-template/pkg/rest-api"
)

func main() {

	rest_api.SPIDER_URL = os.Getenv("SPIDER_URL")

	// Run API Server
	rest_api.ApiServer()

}
