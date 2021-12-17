package rqlite

import (
	"database/sql"
	"testing"
)

func Test_Init(t *testing.T) {
	if len(sql.Drivers()) != 1 || sql.Drivers()[0] != "rqlite" {
		t.Fatal("wrong list of registered Drivers")
	}
}

func Test_BadURL(t *testing.T) {
	db, err := sql.Open("rqlite", "xxhkajh ere")
	if err != nil {
		t.Fatal("Invalid rqlite URL opened OK")
	}

	err = db.Ping()
	if err == nil {
		t.Fatal("expected error from Open or Ping")
	}
}