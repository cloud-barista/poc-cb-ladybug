module github.com/cloud-barista/cb-mcas

go 1.16

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/chartmuseum/helm-push v0.10.1
	github.com/cloud-barista/cb-store v0.4.1
	github.com/go-resty/resty/v2 v2.6.0
	github.com/google/uuid v1.3.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/mittwald/go-helm-client v0.8.3-0.20211026133933-a26889186afc
	github.com/sirupsen/logrus v1.8.1
	github.com/swaggo/echo-swagger v1.1.4
	helm.sh/helm/v3 v3.7.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.2
	github.com/googleapis/gnostic v0.5.5 => github.com/google/gnostic v0.5.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
