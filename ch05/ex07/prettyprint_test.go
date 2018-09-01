package main

import (
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestPrettyOutputCanBeParsed(t *testing.T) {
	url := "https://golang.org"
	if err := outline(url); err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}
	_, err := html.Parse(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}
}
