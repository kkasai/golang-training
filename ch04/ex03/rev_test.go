package main

import (
	"testing"
)

func TestRev(t *testing.T) {
	s := [Size]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&s)

	expected := [Size]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	if s != expected {
		t.Errorf("expected: %d but was actual: %d", expected, s)
	}
}
