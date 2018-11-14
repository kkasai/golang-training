package main

import "testing"

func benchN(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchN(b, 1, PopCount)
}

func BenchmarkPopCount10(b *testing.B) {
	benchN(b, 10, PopCount)
}

func BenchmarkPopCount100(b *testing.B) {
	benchN(b, 100, PopCount)
}

func BenchmarkPopCount1000(b *testing.B) {
	benchN(b, 1000, PopCount)
}

func BenchmarkPopCount10000(b *testing.B) {
	benchN(b, 10000, PopCount)
}

func BenchmarkPopCountFor(b *testing.B) {
	benchN(b, 1, PopCountFor)
}

func BenchmarkPopCountFor10(b *testing.B) {
	benchN(b, 10, PopCountFor)
}

func BenchmarkPopCountFor100(b *testing.B) {
	benchN(b, 100, PopCountFor)
}

func BenchmarkPopCountFor1000(b *testing.B) {
	benchN(b, 1000, PopCountFor)
}

func BenchmarkPopCountFor10000(b *testing.B) {
	benchN(b, 10000, PopCountFor)
}

func BenchmarkPopCount64Shift(b *testing.B) {
	benchN(b, 1, PopCount64Shift)
}

func BenchmarkPopCount64Shift10(b *testing.B) {
	benchN(b, 10, PopCount64Shift)
}

func BenchmarkPopCount64Shift100(b *testing.B) {
	benchN(b, 100, PopCount64Shift)
}

func BenchmarkPopCount64Shift1000(b *testing.B) {
	benchN(b, 1000, PopCount64Shift)
}

func BenchmarkPopCount64Shift10000(b *testing.B) {
	benchN(b, 10000, PopCount64Shift)
}

func BenchmarkPopCountOne(b *testing.B) {
	benchN(b, 1, PopCountOne)
}

func BenchmarkPopCountOne10(b *testing.B) {
	benchN(b, 10, PopCountOne)
}

func BenchmarkPopCountOne100(b *testing.B) {
	benchN(b, 100, PopCountOne)
}

func BenchmarkPopCountOne1000(b *testing.B) {
	benchN(b, 1000, PopCountOne)
}

func BenchmarkPopCountOne10000(b *testing.B) {
	benchN(b, 10000, PopCountOne)
}
