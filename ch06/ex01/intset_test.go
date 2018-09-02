package main

import (
	"testing"
)

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		values   []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 144, 9, 42}, 4},
	} {
		var x IntSet
		for _, v := range tc.values {
			x.Add(v)
		}
		if x.Len() != tc.expected {
			t.Errorf("expected: %d but was actual: %d", tc.expected, x.Len())
		}
	}

	var x IntSet
	if x.Len() != 0 {
		t.Errorf("expected: 0 but was actual: %d", x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet

	x.Remove(1000)
	if x.Len() != 0 {
		t.Errorf("expected: 0 but was actual: %d", x.Len())
	}

	const max = 100
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	for i := 0; i < max; i++ {
		x.Remove(i)
		if x.Has(i) {
			t.Errorf("x.Has(%d) is true, but want false", i)
			continue
		}
		if x.Len() != (max - i - 1) {
			t.Errorf("expected: %d but was actual: %d", max-i-1, x.Len())
		}
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	const max = 100
	for i := 0; i < max; i++ {
		x.Add(i)
	}
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("expected: 0 but was actual: %d", x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	const max = 100
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	c := x.Copy()
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("expected: 0 but was actual: %d", x.Len())
	}

	for i := 0; i < max; i++ {
		if x.Has(i) {
			t.Errorf("c.Has(%d) is true, but want false", i)
		}
		if !c.Has(i) {
			t.Errorf("c.Has(%d) is false, but want true", i)
		}
	}
}
