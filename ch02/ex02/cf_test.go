package main

import (
	"bytes"
	"os"
	"testing"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestEcho(t *testing.T) {
	os.Args = []string{"cmd", "-t", "10", "-d", "10", "-w", "10"}
	main()

	s := "10℉ = -12.222222222222221℃, 10℃ = 50℉\n" +
		"10m = 32.808398950131235ft, 10ft = 3.048m\n" +
		"10lb = 4.535923700000001kg, 10kg = 22.046226218487757lb\n"
	if buffer.String() != s {
		t.Errorf("expected: %s but was actual: %s", s, buffer.String())
	}
}
