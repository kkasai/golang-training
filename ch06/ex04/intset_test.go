package main

import "testing"

func TestEnums(t *testing.T) {
	for _, tc := range []struct {
		t []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 10, 100, 1000, 10000}},
	} {
		var x IntSet
		x.AddAll(tc.t...)

		elems := x.Elems()
		if len(elems) != len(tc.t) {
			t.Errorf("expected: %d but was actual: %d", len(tc.t), len(elems))
			t.Errorf("elems: %v, t: %v", elems, tc.t)
		}

		for _, value := range elems {
			if !x.Has(value) {
				t.Errorf("expected: ture but was actual: false. Value: %d", value)
			}
		}
	}
}
