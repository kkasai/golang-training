package main

import (
	"crypto/sha256"
	"testing"
)

var data = []struct {
	d1       []byte
	d2       []byte
	expected int
}{
	{[]byte("x"), []byte("x"), 0},
	{[]byte("x"), []byte("X"), 125},
}

func TestSha(t *testing.T) {
	for _, d := range data {
		c1 := sha256.Sum256(d.d1)
		c2 := sha256.Sum256(d.d2)
		result := sha(c1, c2)
		if result != d.expected {
			t.Errorf("expected: %d but was actual: %d", d.expected, result)
		}
	}
}
