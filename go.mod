module github.com/cloud-barista/cb-mcas

go 1.16

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/chartmuseum/helm-push v0.9.0
	github.com/cloud-barista/cb-store v0.4.1
	github.com/go-resty/resty/v2 v2.6.0
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/labstack/echo/v4 v4.2.1
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/swaggo/echo-swagger v1.1.3
	github.com/vmware-labs/yaml-jsonpath v0.3.2
	go.uber.org/zap v1.17.0 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/cli-runtime v0.22.1 // indirect
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.2
	//	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.14
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
