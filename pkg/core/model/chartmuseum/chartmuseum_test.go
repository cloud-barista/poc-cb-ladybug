package chartmuseum

import (
	"testing"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/config"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/utils/lang"
)

var (
	repoName     string = "local"
	chartName    string = "chartmuseum"
	chartVersion string = "3.2.0"
)

func TestMain(t *testing.T) {
	config.Setup()
}

func TestGetAllCharts(t *testing.T) {
	cm := NewChartmuseum(repoName)
	chartsmap, err := cm.GetAllCharts()
	if err != nil {
		t.Fatalf("Chartmuseum.GetAllCharts error - cause=%v : %s", err, lang.GetFuncName())
	}

	printChartsMap(t, chartsmap)
}

func TestChartCRD(t *testing.T) {
	// upload a chart
	cm := NewChartmuseum(repoName)
	err := cm.UploadChart("../../../../scripts/chartmuseum-3.2.0.tgz")
	if err != nil {
		t.Fatalf("Chartmuseum.UploadChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	// verify upload a chart
	chart, err := cm.GetChart(chartName, chartVersion)
	if err != nil {
		t.Fatalf("Chartmuseum.GetChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	chartsmap, err := cm.GetAllCharts()
	if err != nil {
		t.Fatalf("Chartmuseum.GetAllCharts error - cause=%v : %s", err, lang.GetFuncName())
	}

	printChartsMap(t, chartsmap)

	chartList, err := cm.GetAllVersionsOfChart(chartName)
	if err != nil {
		t.Fatalf("Chartmuseum.GetAllVersionsOfChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	if chartList != nil {
		for _, ch := range *chartList {
			t.Log("\t", "name:", ch.Name, ", version:", ch.Version)
		}
	}

	chart, err = cm.GetChart(chartName, chartVersion)
	if err != nil {
		t.Fatalf("Chartmuseum.GetChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	if chart != nil {
		t.Log("\t", "name:", chart.Name, ", version:", chart.Version)
	}

	err = cm.DeleteChart(chartName, chartVersion)
	if err != nil {
		t.Fatalf("Chartmuseum.DeleteChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	chart, err = cm.GetChart(chartName, chartVersion)
	if err != nil {
		t.Fatalf("Chartmuseum.GetChart error - cause=%v : %s", err, lang.GetFuncName())
	}

	if chart != nil {
		t.Logf("Fail to delete a chart (%s-%s)", chartName, chartVersion)
	}
}

func TestDeleteNonexistingChart(t *testing.T) {
	ver := "3.0.0"
	cm := NewChartmuseum(repoName)
	err := cm.DeleteChart(chartName, ver)
	if err != nil {
		t.Fatalf("Chartmuseum.DeleteChart error - cause=%v : %s", err, lang.GetFuncName())
	}
}

func printChartsMap(t *testing.T, chartsmap *map[string][]CmChart) {
	for key, cm := range *chartsmap {
		t.Log("Key:", key)
		for _, c := range cm {
			t.Log("\t", " name: ", c.Name, ", version: ", c.Version)
		}
	}
}
