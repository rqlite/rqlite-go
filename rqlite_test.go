package rqlite

import (
	"database/sql"
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test_Init(t *testing.T) {
	if len(sql.Drivers()) != 1 || sql.Drivers()[0] != "rqlite" {
		t.Fatal("wrong list of registered Drivers")
	}
}

func Test_SimpleOpen(t *testing.T) {
	_, err := sql.Open("rqlite", "http://example.com")
	if err != nil {
		t.Fatal("failed to perform simple DB Open")
	}
}

func Test_SimplePingFail(t *testing.T) {
	db, err := sql.Open("rqlite", "http://example.com")
	if err != nil {
		t.Fatal("failed to perform simple DB Open")
	}
	if db.Ping() == nil {
		t.Fatal("successfully pinged non-existent DB")
	}
}

func Test_SimplePingOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	db, err := sql.Open("rqlite", ts.URL)
	if err != nil {
		t.Fatal("failed to perform simple DB Open")
	}
	if db.Ping() != nil {
		t.Fatal("failed to ping mock DB")
	}
}