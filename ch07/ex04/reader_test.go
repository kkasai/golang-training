package main

import (
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

func TestNewReader(t *testing.T) {
	s := "hello"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(newReader(s))
	if n != int64(len(s)) || err != nil {
		t.Logf("n=%d err=%s", n, err)
		t.Fail()
	}
	if b.String() != s {
		t.Logf(`"%s" != "%s"`, b.String(), s)
	}
}

func TestNewReaderWithHTML(t *testing.T) {
	s := "<html><body><p>hello</p></body></html>"
	_, err := html.Parse(newReader(s))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
