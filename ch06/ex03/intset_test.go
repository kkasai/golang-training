package main

import "testing"

func TestIntersectWith(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{4, 5}},
		{[]int{1, 10, 100, 1000}, []int{1, 10}, []int{1, 10}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.IntersectWith(&y)
		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("expected: ture but was actual: false. Value: %d", value)
			}
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3}},
		{[]int{1, 10, 100, 1000}, []int{1, 10}, []int{100, 1000}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.DifferenceWith(&y)
		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("expected: ture but was actual: false. Value: %d", value)
			}
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	for _, tc := range []struct {
		t        []int
		s        []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3, 6, 7, 8}},
		{[]int{1, 10, 100, 1000}, []int{1, 10}, []int{100, 1000}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(tc.t...)
		y.AddAll(tc.s...)

		x.SymmetricDifference(&y)
		for _, value := range tc.expected {
			if !x.Has(value) {
				t.Errorf("expected: ture but was actual: false. Value: %d", value)
			}
		}
	}
}
