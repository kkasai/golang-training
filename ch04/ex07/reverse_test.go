package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			"abcdef",
			"fedcba"},
		{
			"あいうえお",
			"おえういあ"},
	}

	for _, d := range data {
		comp := reverse([]byte(d.s))
		result := string(comp)
		if result != d.expected {
			t.Errorf(`expected: %s but was actual: %s"`,
				d.expected, result)
		}
	}
}
