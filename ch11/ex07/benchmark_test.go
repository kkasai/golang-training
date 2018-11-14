package intset

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set IntSet
		for j := 0; j < 500; j++ {
			set.Add(rand.Intn(math.MaxInt16))
		}
	}
}

func BenchmarkMapIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewMapIntSet()
		for j := 0; j < 500; j++ {
			set.Add(rand.Intn(math.MaxInt16))
		}
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x IntSet
		var y IntSet
		for j := 0; j < 500; j++ {
			x.Add(rand.Intn(math.MaxInt16))
			y.Add(rand.Intn(math.MaxInt16))
		}
		x.UnionWith(&y)
	}
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := NewMapIntSet()
		y := NewMapIntSet()
		for j := 0; j < 500; j++ {
			x.Add(rand.Intn(math.MaxInt16))
			y.Add(rand.Intn(math.MaxInt16))
		}
		x.UnionWith(y)
	}
}

func BenchmarkIntSet_AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var set IntSet
		ints := make([]int, 500)
		for j := 0; j < 500; j++ {
			ints[j] = rand.Intn(math.MaxInt16)
		}
		set.AddAll(ints...)
	}
}

func BenchmarkMapIntSet_AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := NewMapIntSet()
		ints := make([]int, 500)
		for j := 0; j < 500; j++ {
			ints[j] = rand.Intn(math.MaxInt16)
		}
		set.AddAll(ints...)
	}
}
