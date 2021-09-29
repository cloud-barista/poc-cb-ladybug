package common

import (
	"time"

	cbstore "github.com/cloud-barista/cb-store"
	cbsconfig "github.com/cloud-barista/cb-store/config"
	cbsinterface "github.com/cloud-barista/cb-store/interfaces"
	"github.com/sirupsen/logrus"
)

// CB-Store
var CBLog *logrus.Logger
var CBStore cbsinterface.Store

var StartTime string

func init() {
	CBLog = cbsconfig.Cblogger
	CBStore = cbstore.GetStore()

	StartTime = time.Now().Format("2006.01.02 15:04:05 Mon")
}
