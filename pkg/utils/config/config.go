package config

import (
	"flag"
	"os"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/lang"
)

type conf struct {
	RunMode             *string
	AppRootPath         *string
	ListenAddress       *string
	BasePath            *string
	SpiderCallMethod    *string
	TumblebugCallMethod *string
	McksCallMethod      *string
	SpiderUrl           *string
	TumblebugUrl        *string
	McksUrl             *string
	Username            *string
	Password            *string
	LoglevelHTTP        *bool
}

var Config = &conf{}

func Setup() {

	//	var logLevel *string

	Config.AppRootPath = flag.String("app-root", lang.NVL(os.Getenv("APP_ROOT"), ""), "application root path")
	Config.ListenAddress = flag.String("listen-address", lang.NVL(os.Getenv("LISTEN_ADDRESS"), ":1592"), "ladybug listen address(IP:port)")
	Config.BasePath = flag.String("base-path", lang.NVL(os.Getenv("BASE_PATH"), "/ladybug"), "ladybug base path")
	Config.SpiderCallMethod = flag.String("spider-call-method", lang.NVL(os.Getenv("SPIDER_CALL_METHOD"), "REST"), "Method of calling CB-Spider (REST/gRPC)")
	Config.TumblebugCallMethod = flag.String("tumblebug-call-method", lang.NVL(os.Getenv("TUMBLEBUG_CALL_METHOD"), "REST"), "Method of calling CB-Tumblebug (REST/gRPC)")
	Config.McksCallMethod = flag.String("mcks-call-method", lang.NVL(os.Getenv("MCKS_CALL_METHOD"), "REST"), "Method of calling CB-MCKS(REST/gRPC)")
	Config.SpiderUrl = flag.String("spider-url", lang.NVL(os.Getenv("SPIDER_URL"), "http://localhost:1024/spider"), "cb-spider service end-point url")
	Config.TumblebugUrl = flag.String("tumblebug-url", lang.NVL(os.Getenv("TUMBLEBUG_URL"), "http://localhost:1323/tumblebug"), "cb-tumblebug service end-point url")
	Config.McksUrl = flag.String("mcks-url", lang.NVL(os.Getenv("MCKS_URL"), "http://localhost:8080/ladybug"), "cb-mcks service end-point url")
	Config.Username = flag.String("basic-auth-username", lang.NVL(os.Getenv("BASIC_AUTH_USERNAME"), "default"), "rest-api basic auth usernmae")
	Config.Password = flag.String("basic-auth-password", lang.NVL(os.Getenv("BASIC_AUTH_PASSWORD"), "default"), "rest-api basic auth password")
	//	logLevel = flag.String("log-level", lang.NVL(os.Getenv("LOG_LEVEL"), "info"), "The log level")
	//	Config.LoglevelHTTP = flag.Bool("log-http", os.Getenv("LOG_HTTP") == "true", "The logging http data")

	flag.Parse()

	// CBLog
	//	common.CBLog.SetFormatter(&logrus.TextFormatter{})
	/*
		common.CBLog.SetOutput(os.Stderr)

		level, err := logrus.ParseLevel(*logLevel)
		if err != nil {
			common.CBLog.Fatal(err)
		} else if level != common.CBLog.GetLevel() {
			common.CBLog.SetLevel(level)
		} else {
			common.CBLog.SetLevel(logrus.InfoLevel)
		}
	*/
	// app root path
	if len(*Config.AppRootPath) == 0 {
		if pwd, err := os.Getwd(); err == nil {
			Config.AppRootPath = &pwd
		}
	}

	common.CBLog.Infof("app-root: ", *Config.AppRootPath)
	common.CBLog.Infof("listen-address: ", *Config.ListenAddress)
	common.CBLog.Infof("base-path: ", *Config.BasePath)
	common.CBLog.Infof("spder-call-method: ", *Config.SpiderCallMethod)
	common.CBLog.Infof("tumblebug-call-method: ", *Config.TumblebugCallMethod)
	common.CBLog.Infof("mcks-call-method: ", *Config.McksCallMethod)
	common.CBLog.Infof("spider-url: ", *Config.SpiderUrl)
	common.CBLog.Infof("tumblebug-url: ", *Config.TumblebugUrl)
	common.CBLog.Infof("mcks-url: ", *Config.McksUrl)
	common.CBLog.Infof("basic-auth-username: ", *Config.Username)
	common.CBLog.Infof("basic-auth-password: ", *Config.Password)
	//	common.CBLog.Infof("log-level: ", *logLevel)
	//	common.CBLog.Infof("log-http: ", *Config.LoglevelHTTP)
}
