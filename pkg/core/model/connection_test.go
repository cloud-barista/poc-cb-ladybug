package model

import (
	"testing"
)

func TestConnectionCRD(t *testing.T) {

	namespace := "namespace-1"
	connectionName := "connection-1"

	// insert
	connection := NewConnection(namespace, connectionName)
	err := connection.Insert()
	if err != nil {
		t.Fatalf("Connection.Insert error - connection.Insert() (cause=%v)", err)
	}

	// verify insert
	connection = NewConnection(namespace, connectionName)
	err = connection.Select()
	if err != nil {
		t.Fatalf("Connection.Insert error - connection.Select() (cause=%v)", err)
	}

	// delete
	connection = NewConnection(namespace, connectionName)
	err = connection.Delete()
	if err != nil {
		t.Fatalf("Connection.Delete error - connection.Delete() (cause=%v)", err)
	}

	// verify delete
	connection = NewConnection(namespace, connectionName)
	err = connection.Select()
	if err == nil {
		t.Fatalf("Connection.Delete exist data, not deleted")
	}

}

func TestConnectionListSelect(t *testing.T) {
	namespace := "namespace-1"
	connectionName := "connection-1"

	// list
	connectionList := NewConnectionList(namespace)
	err := connectionList.SelectList()
	if err != nil {
		t.Fatalf("error (cause=%v)", err)
	}

	// delete all listed items
	for _, cfg := range connectionList.Items {
		err = cfg.Delete()
		if err != nil {
			t.Fatalf("Connection.Delete error - connection.Delete() (cause=%v)", err)
		}
	}

	// insert
	connection1 := NewConnection(namespace, connectionName)
	err = connection1.Insert()
	if err != nil {
		t.Fatalf("connection.Insert error (cause=%v)", err)
	}

	// insert
	connectionName = "connection-2"
	connection2 := NewConnection(namespace, connectionName)
	err = connection2.Insert()
	if err != nil {
		t.Fatalf("connection.Insert error (cause=%v)", err)
	}

	// list
	connectionList = NewConnectionList(namespace)
	err = connectionList.SelectList()
	if err != nil {
		t.Fatalf("error (cause=%v)", err)
	}

	if len(connectionList.Items) != 2 {
		t.Fatalf("missmatched rows (count=%v)", len(connectionList.Items))
	}

	// delete
	err = connection1.Delete()
	if err != nil {
		t.Fatalf("Connection.Delete error - connection.Delete() (cause=%v)", err)
	}

	err = connection2.Delete()
	if err != nil {
		t.Fatalf("Connection.Delete error - connection.Delete() (cause=%v)", err)
	}

	// list
	connectionList = NewConnectionList(namespace)
	err = connectionList.SelectList()
	if err != nil {
		t.Fatalf("error (cause=%v)", err)
	}

	if len(connectionList.Items) != 0 {
		t.Fatalf("missmatched rows (count=%v)", len(connectionList.Items))
	}
}
