package main

import "testing"

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		values   []int
		expected string
	}{
		{[]int{1}, "{1}"},
		{[]int{1, 144, 9, 42}, "{1 9 42 144}"},
	} {
		var x IntSet
		for _, v := range tc.values {
			x.Add(v)
		}
		if x.String() != tc.expected {
			t.Errorf("expected: %s but was actual: %s", tc.expected, x.String())
		}
	}
}
