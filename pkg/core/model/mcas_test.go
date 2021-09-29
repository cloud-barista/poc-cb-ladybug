package model

import (
	"testing"
)

func TestMcasCRD(t *testing.T) {
	var status string

	namespace := "namespace-1"

	// init
	mcas := NewMcas(namespace)
	err := mcas.Init()
	if err != nil {
		t.Fatalf("Mcas.Init error - Mcas.Init() (cause=%v)", err)
	}

	// verify init
	mcas = NewMcas(namespace)
	status, err = mcas.GetStatus()
	if err != nil {
		t.Fatalf("Mcas.GetStatus error - Mcas.GetStatus() (cause=%v)", err)
	}

	if status != STATUS_MCAS_ENABLED {
		t.Fatalf("Mcas.Init is failed: status=%s", status)
	}

	// disable
	mcas = NewMcas(namespace)
	err = mcas.Disable()
	if err != nil {
		t.Fatalf("Mcas.Disable error - Mcas.Disable() (cause=%v)", err)
	}

	// verify disable
	mcas = NewMcas(namespace)
	status, err = mcas.GetStatus()
	if err != nil {
		t.Fatalf("Mcas.GetStatus error - Mcas.GetStatus() (cause=%v)", err)
	}

	if status != STATUS_MCAS_DISABLED {
		t.Fatalf("Mcas.Disable is failed")
	}

	// enable
	mcas = NewMcas(namespace)
	err = mcas.Enable()
	if err != nil {
		t.Fatalf("Mcas.Enable error - Mcas.Enable() (cause=%v)", err)
	}

	// verify disable
	mcas = NewMcas(namespace)
	status, err = mcas.GetStatus()
	if err != nil {
		t.Fatalf("Mcas.GetStatus error - Mcas.GetStatus() (cause=%v)", err)
	}

	if status != STATUS_MCAS_ENABLED {
		t.Fatalf("Mcas.Enable is failed")
	}

	// delete
	mcas = NewMcas(namespace)
	err = mcas.Delete()
	if err != nil {
		t.Fatalf("Mcas.Delete error - Mcas.Delete() (cause=%v)", err)
	}

	// verify delete
	mcas = NewMcas(namespace)
	status, err = mcas.GetStatus()
	if err == nil {
		t.Fatalf("Mcas.Delete exist data, not deleted")
	}
}
