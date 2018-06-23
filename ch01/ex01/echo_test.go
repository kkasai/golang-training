package main

import (
	"testing"
	"os"
	"bytes"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestEcho(t *testing.T) {
	os.Args = []string{"cmd", "a", "b", "c", "d"}
	main()

	if buffer.String() != "cmd a b c d\n" {
		t.Errorf("expected: cmd a b c d but was actual: %s", buffer.String())
	}
}
