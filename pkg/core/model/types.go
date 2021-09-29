package model

const (
	KIND_STATUS            = "Status"
	KIND_MCAS              = "MCAService"
	KIND_CONNECTION        = "Connection"
	KIND_CONNECTION_LIST   = "ConnectionList"
	KIND_PACKAGE           = "Package"
	KIND_PACKAGE_LIST      = "PackageList"
	KIND_APP_INSTANCE      = "AppInstance"
	KIND_APP_INSTANCE_LIST = "AppInstanceList"
	KIND_CLUSTER           = "Cluster"
	KIND_CLUSTER_LIST      = "ClusterList"
)

type Model struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
type ListModel struct {
	Kind string `json:"kind"`
}
