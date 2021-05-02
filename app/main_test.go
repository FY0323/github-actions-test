package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(handler))
	defer testserver.Close()

	res, err := http.Get(testserver.URL)
	defer res.Body.Close()

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if res.StatusCode != 200 {
		t.Error("response code is not 200")
	}
}
