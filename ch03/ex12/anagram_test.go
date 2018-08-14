package main

import "testing"

var data = []struct {
	w1       string
	w2       string
	expected bool
}{
	{"Hello", "Hello", true},
	{"Christmas", "trims cash", true},
	{"ぶたがすわらん", "すがわらぶんた", true},
	{"aaaaa", "aaaabbb", false},
}

func TestAnagrams(t *testing.T) {
	for _, d := range data {
		result := anagram(d.w1, d.w2)
		if result != d.expected {
			t.Errorf("expected: %v but was actual: %v, w1: %s, w2: %s", d.expected, result, d.w1, d.w2)
		}
	}
}
