package main

import "testing"

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll()
	if x.Len() != 0 {
		t.Errorf("expected: 0 but was actual: %d", x.Len())
	}

	x.AddAll(1, 2, 3, 4, 5)
	if x.Len() != 5 {
		t.Errorf("expected: 5 but was actual: %d", x.Len())
	}

	for i := 1; i <= 5; i++ {
		if !x.Has(i) {
			t.Errorf("x.Has(%d) is false, but want true", i)
		}
	}
}
