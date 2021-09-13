package tumblebug

import (
	"testing"

	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

func TestMain(t *testing.T) {
	config.Setup()
}

func TestListConnConfig(t *testing.T) {
	tb := NewTumblebug()
	connConfigList, err := tb.ListConnConfig()
	if err != nil {
		t.Fatalf("Tumblebug.ListConnConfig error - cause=%v : %s", err, lang.GetFuncName())
	}

	printConnConfigList(t, connConfigList)
}

func printConnConfigList(t *testing.T, list *TbConnConfigList) {
	t.Log("connenctionconfig:")
	for i, cc := range list.ConnectionConfig {
		t.Logf("[%d]", i)
		t.Log("\tconfigName:", cc.ConfigName)
		t.Log("\tcredentialName:", cc.CredentialName)
		t.Log("\tdriverName:", cc.DriverName)
		t.Log("\tproviderName:", cc.ProviderName)
		t.Log("\tregionName:", cc.RegionName)
	}
}
