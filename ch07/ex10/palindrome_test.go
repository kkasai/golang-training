package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for _, test := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"abcdefedcba", true},
		{"hello world", false},
		{"sator arepo tenet opera rotas", true},
		{"palindrome", false},
	} {
		if got := IsPalindrome(PalindromeString(test.input)); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
