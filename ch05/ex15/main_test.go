package main

import "testing"

func TestMax(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		valid    bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3, 4, 5}, true, 5},
		{[]int{1, 2, 5, 3, 4}, true, 5},
	} {
		m, err := max(test.vals...)
		if !test.valid {
			if err == nil {
				t.Errorf("err != nil expected for %v\n", test.vals)
			}
			continue
		}
		if m != test.expected {
			t.Errorf("expected: %d but was actual: %d\n", test.expected, m)
		}
	}
}

func TestMax2(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 5, 3, 4}, 5},
	} {
		m := max2(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("expected: %d but was actual: %d\n", test.expected, m)
		}
	}
}

func TestMin(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		valid    bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3, 4, 5}, true, 1},
		{[]int{5, 2, 1, 3, 4}, true, 1},
	} {
		m, err := min(test.vals...)
		if !test.valid {
			if err == nil {
				t.Errorf("err != nil expected for %v\n", test.vals)
			}
			continue
		}
		if m != test.expected {
			t.Errorf("expected: %d but was actual: %d\n", test.expected, m)
		}
	}
}

func TestMin2(t *testing.T) {
	for _, test := range []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 2, 1, 3, 4}, 1},
	} {
		m := min2(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("expected: %d but was actual: %d\n", test.expected, m)
		}
	}
}
