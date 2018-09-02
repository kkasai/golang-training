package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	toLower := func(s string) string { return strings.ToLower(s) }

	for _, test := range []struct {
		s        string
		f        func(string) string
		expected string
	}{
		{"$hello world", toLower, "hello world"},
		{"$FOO Bar $FuGa", toLower, "foo Bar fuga"},
	} {
		result := expand(test.s, test.f)
		if result != test.expected {
			t.Errorf("expected: %s but was actual: %s\n", test.expected, result)
		}
	}
}
