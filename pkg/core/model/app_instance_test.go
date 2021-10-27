package model

import (
	"testing"

	"github.com/cloud-barista/cb-ladybug/pkg/utils/config"
)

func TestMain(t *testing.T) {
	config.Setup()
}

/*
func TestAppInstanceCRD(t *testing.T) {

	namespace := "namespace-1"
	appInstanceName := "app_instance-1"
	appPkgName := "app_package-1"
	version := "1.0.0"

	// insert
	appInstance := NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err := appInstance.Insert()
	if err != nil {
		t.Fatalf("AppInstance.Insert error - pkg.Insert() (cause=%v)", err)
	}

	// verify insert
	appInstance = NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err = appInstance.Select()
	if err != nil {
		t.Fatalf("AppInstance.Insert error - pkg.Select() (cause=%v)", err)
	}

	// delete
	appInstance = NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err = appInstance.Delete()
	if err != nil {
		t.Fatalf("AppInstance.Delete error - pkg.Delete() (cause=%v)", err)
	}

	// verify delete
	appInstance = NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err = appInstance.Select()
	if err == nil {
		t.Fatalf("AppInstance.Delete exist data, not deleted")
	}

}

func TestAppInstanceListSelect(t *testing.T) {
	namespace := "namespace-1"
	appInstanceName := "app_instance-1"
	appPkgName := "app_package-1"
	version := "1.0.0"

	// list
	appInstanceList := NewAppInstanceList(namespace)
	err := appInstanceList.SelectList()
	if err != nil {
		t.Fatalf("error (cause=%v)", err)
	}

	// delete all listed items
	for _, ai := range appInstanceList.Items {
		err = ai.Delete()
		if err != nil {
			t.Fatalf("AppInstance.Delete error (cause=%v)", err)
		}
	}

	// insert
	appInstance1 := NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err = appInstance1.Insert()
	if err != nil {
		t.Fatalf("AppInstance.Insert error (cause=%v)", err)
	}

	// insert
	appInstanceName = "app_instance-2"
	appInstance2 := NewAppInstance(namespace, appInstanceName, appPkgName, version)
	err = appInstance2.Insert()
	if err != nil {
		t.Fatalf("AppInstance.Insert error (cause=%v)", err)
	}

	// list
	appInstanceList = NewAppInstanceList(namespace)
	err = appInstanceList.SelectList()
	if err != nil {
		t.Fatalf("AppInstanceList.SelectList error (cause=%v)", err)
	}

	if len(appInstanceList.Items) != 2 {
		t.Fatalf("missmatched rows (count=%v)", len(appInstanceList.Items))
	}

	err = appInstance1.Delete()
	if err != nil {
		t.Fatalf("AppInstance.Delete error (cause=%v)", err)
	}

	err = appInstance2.Delete()
	if err != nil {
		t.Fatalf("AppInstance.Delete error (cause=%v)", err)
	}

	// list
	appInstanceList = NewAppInstanceList(namespace)
	err = appInstanceList.SelectList()
	if err != nil {
		t.Fatalf("error (cause=%v)", err)
	}

	if len(appInstanceList.Items) != 0 {
		t.Fatalf("missmatched rows (count=%v)", len(appInstanceList.Items))
	}
}
*/
