package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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

	if buffer.String() != "200 OK : Fetch Test" {
		t.Errorf("expected: Fetch Test but was actual: %s", buffer.String())
	}
}

func initBuffer() {
	buffer = &bytes.Buffer{}
	writer = buffer
}
