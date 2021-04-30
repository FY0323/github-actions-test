package main

import (
	"io/ioutil"
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

	hello, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if string(hello) != "hello world" {
		t.Error("response is not hello world")
	}
}
