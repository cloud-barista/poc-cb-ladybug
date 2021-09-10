module github.com/cloud-barista/cb-mcas

go 1.15

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/chartmuseum/helm-push v0.9.0
	github.com/cloud-barista/cb-store v0.4.1
	github.com/go-resty/resty/v2 v2.6.0
	github.com/google/uuid v1.2.0
	github.com/labstack/echo/v4 v4.2.1
	github.com/sirupsen/logrus v1.8.1
	github.com/swaggo/echo-swagger v1.1.3
	github.com/vmware-labs/yaml-jsonpath v0.3.2
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.2
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.8
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
