package main

import (
	"testing"
)

func TestRemove(t *testing.T) {
	data := []struct {
		s        []string
		expected []string
	}{
		{
			[]string{"aaa"},
			[]string{"aaa"}},
		{
			[]string{"aaa", "bbb", "bbb"},
			[]string{"aaa", "bbb"}},
		{
			[]string{"aaa", "bbb", "bbb", "ccc"},
			[]string{"aaa", "bbb", "ccc"}},
	}

	for _, d := range data {
		result := remove(d.s)
		if len(result) != len(d.expected) {
			t.Errorf("expected length: %d but was actual: %d",
				len(result), len(d.expected))
		}
		for i := 0; i < len(d.expected); i++ {
			if result[i] != d.expected[i] &&
				d.s[i] != d.expected[i] {
				t.Errorf(`expected: %s but was actual: %s"`,
					result[i], d.expected[i])
			}
		}
	}
}
