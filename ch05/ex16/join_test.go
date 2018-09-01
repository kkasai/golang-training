package main

import (
	"testing"
)

var testCases = []struct {
	a        []string
	sep      string
	expected string
}{
	{[]string{"abc"}, " ", "abc"},
	{[]string{"abc", "def", "ghi"}, " ", "abc def ghi"},
}

func TestJoin(t *testing.T) {
	for _, tc := range testCases {
		result, err := join(tc.sep, tc.a...)
		if err != nil {
			t.Errorf("%s", err)
		}
		if result != tc.expected {
			t.Errorf("Join(%s, %v) = %s, want %s",
				tc.sep, tc.a, result, tc.expected)
		}
	}
}
