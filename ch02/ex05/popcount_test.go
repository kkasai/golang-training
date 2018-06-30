package main

import "testing"

func TestPopCountOne(t *testing.T) {
	c := PopCountOne(uint64(200))
	if c != 3 {
		t.Errorf("expected: 3 but was actual: %d", c)
	}
}

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(100))
	}
}

func BenchmarkPopCount64Shift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64Shift(uint64(100))
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFor(uint64(100))
	}
}

func BenchmarkPopCountOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOne(uint64(100))
	}
}
