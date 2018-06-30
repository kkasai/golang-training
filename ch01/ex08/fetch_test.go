package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var buffer *bytes.Buffer
var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fetch Test")
})

func TestFetch(t *testing.T) {
	initBuffer()

	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	os.Args = []string{"cmd", ts.URL}
	main()

	if buffer.String() != "Fetch Test" {
		t.Errorf("expected: Fetch Test but was actual: %s", buffer.String())
	}
}

func TestNoSchemeFetch(t *testing.T) {
	initBuffer()

	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	serverUrl, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("parse err: %s", err)
	}
	os.Args = []string{"cmd", serverUrl.Host}
	main()

	if buffer.String() != "Fetch Test" {
		t.Errorf("expected: Fetch Test but was actual: %s", buffer.String())
	}
}

func initBuffer() {
	buffer = &bytes.Buffer{}
	writer = buffer
}
