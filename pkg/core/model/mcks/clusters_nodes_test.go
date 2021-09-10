package mcks

import (
	"testing"

	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

const (
	MCKS_STATUS_COMPLETED = "completed"
)

var (
	namespace   string = "namespace-1"
	clusterName string = "cluster-1"
)

//
// Note:
// you should test with 20min timeout,
// because it takes about 10min to create a cluster
// $ go test -v -timeout 20m
//
func TestMain(t *testing.T) {
	config.Setup()
}

func TestCreateGetCluster(t *testing.T) {
	// create a cluster
	mcks := NewMcks(namespace)
	clusterReq := makeMcksClusterReq(clusterName)
	cluster, err := mcks.CreateCluster(*clusterReq)
	if err != nil {
		mcks.DeleteCluster(clusterName)
		t.Fatalf("Mcks.CreateCluster error - cause=%v : %s", err, lang.GetFuncName())
	}

	if cluster.Status != MCKS_STATUS_COMPLETED {
		mcks.DeleteCluster(clusterName)
		t.Fatalf("The cluster creation is not completed - cause=%v : %s", err, lang.GetFuncName())
	}

	printMcksCluster(t, cluster)

	// verify the created cluster
	clusterNew, err := mcks.GetCluster(clusterName)
	if err != nil {
		t.Fatalf("Mcks.GetCluster error - cause=%v : %s", err, lang.GetFuncName())
	}

	if clusterNew == nil {
		t.Fatalf("The cluster(%s) could not found.", clusterName)
	}

	printMcksCluster(t, clusterNew)
}

func TestNodeCRD(t *testing.T) {
	// check the cluster
	mcks := NewMcks(namespace)
	cluster, err := mcks.GetCluster(clusterName)
	if err != nil {
		t.Fatalf("Mcks.GetCluster error - cause=%v : %s", err, lang.GetFuncName())
	}

	if cluster == nil {
		t.Fatalf("The cluster(%s) doess not exist", clusterName)
	}

	nodeReq := makeMcksNodeReq()
	nodeList, err := mcks.AddNodes(clusterName, nodeReq)
	if err != nil {
		t.Fatalf("Mcks.AddNodes error - cause=%v : %s", err, lang.GetFuncName())
	}

	printMcksNodeList(t, nodeList)

	// verify the created nodes
	for _, node := range nodeList.Items {
		n, err := mcks.GetNode(clusterName, node.Name)
		if err != nil {
			t.Errorf("Mcks.Get error - cause=%v : %s", err, lang.GetFuncName())
		}

		if n != nil {
			printMcksNode(t, n)
		}

		// delete the node
		_, err = mcks.RemoveNode(clusterName, node.Name)
		if err != nil {
			t.Errorf("Mcks.Remove error - cause=%v : %s", err, lang.GetFuncName())
		}

		n, err = mcks.GetNode(clusterName, node.Name)
		if err != nil {
			t.Errorf("Mcks.Get error - cause=%v : %s", err, lang.GetFuncName())
		}

		if n != nil {
			printMcksNode(t, n)
			t.Errorf("The deletion of the node(%s) is failed.", node.Name)
		}
	}
}

func TestDeleteCluster(t *testing.T) {
	// delete the cluster
	mcks := NewMcks(namespace)
	_, err := mcks.DeleteCluster(clusterName)
	if err != nil {
		t.Fatalf("Mcks.DeleteCluster error - cause=%v : %s", err, lang.GetFuncName())
	}

	// verify deletion of the cluster
	cluster, err := mcks.GetCluster(clusterName)
	if err != nil {
		t.Fatalf("Mcks.GetCluster error - cause=%v : %s", err, lang.GetFuncName())
	}

	if cluster != nil {
		printMcksCluster(t, cluster)
		t.Fatalf("The deletion of the cluster(%s) is failed", clusterName)
	}
}

func makeMcksClusterReq(clusterName string) *McksClusterReq {
	var clusterReq McksClusterReq

	clusterReq.Config.Kubernetes.NetworkCni = "kilo"
	clusterReq.Config.Kubernetes.PodCidr = "10.244.0.0/16"
	clusterReq.Config.Kubernetes.ServiceCidr = "10.96.0.0/12"
	clusterReq.Config.Kubernetes.ServiceDnsDomain = "cluster.local"

	var ncCp McksNodeConfig
	ncCp.Connection = "config-aws-ap-northeast-2"
	ncCp.Count = 1
	ncCp.Spec = "t2.medium"

	clusterReq.ControlPlane = append(clusterReq.ControlPlane, ncCp)

	clusterReq.Name = clusterName

	var ncW McksNodeConfig
	ncW.Connection = "config-aws-ap-northeast-1"
	ncW.Count = 1
	ncW.Spec = "t2.small"

	clusterReq.Worker = append(clusterReq.Worker, ncW)

	ncW.Connection = "config-gcp-asia-northeast3"
	ncW.Count = 1
	ncW.Spec = "n1-standard-2"

	clusterReq.Worker = append(clusterReq.Worker, ncW)

	return &clusterReq
}

func printMcksCluster(t *testing.T, mcksCluster *McksCluster) {
	t.Log("Cluster Namespace: ", mcksCluster.Namespace)
	t.Log("Cluster Name: ", mcksCluster.Name)
	t.Log("\tClusterConfig: ", mcksCluster.ClusterConfig)
	t.Log("\tCpLeader: ", mcksCluster.CpLeader)
	t.Log("\tKind: ", mcksCluster.Kind)
	t.Log("\tNetworkCni: ", mcksCluster.NetworkCni)
	t.Log("\tNodes: ")
	for _, node := range mcksCluster.Nodes {
		t.Log("\t\tCredential: ", node.Credential)
		t.Log("\t\tCsp: ", node.Csp)
		t.Log("\t\tKind: ", node.Kind)
		t.Log("\t\tName: ", node.Name)
		t.Log("\t\tPublicIp: ", node.PublicIp)
		t.Log("\t\tRole: ", node.Role)
		t.Log("\t\tSpec: ", node.Spec)
		t.Log("\t\tUid: ", node.Uid)
	}
	t.Log("\tStatus: ", mcksCluster.Status)
	t.Log("\tUid: ", mcksCluster.Uid)
}

func makeMcksNodeReq() *McksNodeReq {
	var nodeReq McksNodeReq
	var ncW McksNodeConfig

	ncW.Connection = "config-gcp-asia-northeast3"
	ncW.Count = 1
	ncW.Spec = "n1-standard-2"

	nodeReq.Worker = append(nodeReq.Worker, ncW)

	return &nodeReq
}

func printMcksNode(t *testing.T, node *McksNode) {
	t.Log("Node Name: ", node.Name)
	t.Log("\tCredential: ", node.Credential)
	t.Log("\tCsp: ", node.Csp)
	t.Log("\tKind: ", node.Kind)
	t.Log("\tName: ", node.Name)
	t.Log("\tPublicIp: ", node.PublicIp)
	t.Log("\tRole: ", node.Role)
	t.Log("\tSpec: ", node.Spec)
	t.Log("\tUid: ", node.Uid)
}

func printMcksNodeList(t *testing.T, nodeList *McksNodeList) {
	for _, node := range nodeList.Items {
		printMcksNode(t, &node)
	}
}
