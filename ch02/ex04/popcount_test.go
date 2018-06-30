package main

import "testing"

func TestPopCountFor(t *testing.T) {
	c := PopCountFor(uint64(200))
	if c != 3 {
		t.Errorf("expected: 3 but was actual: %d", c)
	}
}

func BenchmarkPopCountFormula(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFormula(uint64(100))
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFor(uint64(100))
	}
}
