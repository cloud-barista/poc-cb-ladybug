package chartmuseum

import (
	"testing"

	"github.com/cloud-barista/cb-mcas/pkg/utils/config"
	"github.com/cloud-barista/cb-mcas/pkg/utils/lang"
)

func TestMain(t *testing.T) {
	config.Setup()
}

func TestChartsCRD(t *testing.T) {
	charts := NewCharts("local", "chartmuseum")
	err := charts.POST("../../../../scripts/chartmuseum-3.2.0.tgz")
	if err != nil {
		t.Fatalf("Charts.POST error - %s (cause=%v)", lang.GetFuncName(), err)
	}

	success, err := charts.GET("")
	if err != nil {
		t.Fatalf("Charts.GET error - %s (cause=%v)", lang.GetFuncName(), err)
	}

	if success == true {
		for _, ch := range charts.RespChartList {
			t.Log("\t", "name:", ch.Name, ", version:", ch.Version)
		}
	}

	success, err = charts.GET("3.2.0")
	if err != nil {
		t.Fatalf("Charts.GET error - %s (cause=%v)", lang.GetFuncName(), err)
	}

	if success == true {
		t.Log("\t", "name:", charts.RespChart.Name, ", version:", charts.RespChart.Version)
	}

	err = charts.DELETE("3.2.0")
	if err != nil {
		t.Fatalf("Charts.DELETE error - %s (cause=%v)", lang.GetFuncName(), err)
	}

	success, err = charts.GET("3.2.0")
	if err != nil {
		t.Fatalf("Charts.GET error - %s (cause=%v)", lang.GetFuncName(), err)
	}

	if success != false {
		t.Log("Fail to delete a chart (chartmuseum-3.2.0)")
	}
}

func TestDeleteNonexistingChart(t *testing.T) {
	charts := NewCharts("local", "chartmuseum")
	err := charts.DELETE("3.0.0")
	if err != nil {
		t.Fatalf("Charts.DELETE error - %s (cause=%v)", lang.GetFuncName(), err)
	}
}
