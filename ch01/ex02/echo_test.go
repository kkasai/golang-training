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

	if buffer.String() != "1 a\n2 b\n3 c\n4 d\n" {
		t.Errorf("expected: 1 a\n2 b\n3 c\n4 d\n but was actual: %s", buffer.String())
	}
}
