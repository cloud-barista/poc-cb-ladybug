package chartmuseum

import (
	"testing"

	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

func TestGetChartsMap(t *testing.T) {
	//config.Setup()

	chartsmap, err := NewChartsMap("local")
	if err != nil {
		t.Fatalf("ChartsMap.NewChartsMap error - %s (cause=%v", lang.GetFuncName(), err)
	}

	b, err := chartsmap.GET()
	if err != nil {
		t.Fatalf("ChartsMap.GET error - %s (cause=%v", lang.GetFuncName(), err)
	}

	if b != true {
		t.Fatalf("ChartsMap.GET error - %s (cause=%v", lang.GetFuncName(), err)
	}

	for key, charts := range chartsmap.RespChartsMap {
		t.Log("Key:", key)
		for _, c := range charts {
			t.Log("\t", " name: ", c.Name, ", version: ", c.Version)
		}
	}
}
