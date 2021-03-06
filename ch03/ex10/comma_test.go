package main

import (
	"testing"
)

var data = []struct {
	input    string
	expected string
}{
	{"123", "123"},
	{"123456", "123,456"},
	{"1234567", "1,234,567"},
}

func TestComma(t *testing.T) {
	for _, d := range data {
		result := comma(d.input)
		if result != d.expected {
			t.Errorf("expected: %s but was actual: %s", d.expected, result)
		}
	}
}
