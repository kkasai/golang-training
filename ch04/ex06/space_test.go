package main

import (
	"testing"
)

func TestCompressionSpaces(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			"aaa",
			"aaa"},
		{
			"aaa  bbb",
			"aaa bbb"},
		{
			"aaa  bbb  ccc",
			"aaa bbb ccc"},
	}

	for _, d := range data {
		comp := compressionSpaces([]byte(d.s))
		result := string(comp)
		if result != d.expected {
			t.Errorf(`expected: %s but was actual: %s"`,
				d.expected, result)
		}
	}
}
