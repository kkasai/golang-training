package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	y := NewMapIntSet()

	sets := []int{1, 144, 9, 9, 42}

	for _, v := range sets {
		x.Add(v)
		y.Add(v)
	}

	for i := 0; i < 200; i++ {
		if x.Has(i) == y.Has(i) {
			continue
		}
		t.Errorf("x.Has(%d) is %t and y[%d] is %t\n", i, x.Has(i), i, y.Has(i))
	}
}

func TestUnionWith(t *testing.T) {
	sets1 := []int{1, 2, 3, 4, 5}
	sets2 := []int{3, 5, 7}
	var x IntSet
	var y IntSet
	x.AddAll(sets1...)
	y.AddAll(sets2...)

	x.UnionWith(&y)

	a := NewMapIntSet()
	b := NewMapIntSet()
	a.AddAll(sets1...)
	b.AddAll(sets2...)

	a.UnionWith(b)

	for i := 0; i < 50; i++ {
		if x.Has(i) == a.Has(i) {
			continue
		}
		t.Errorf("x.Has(%d) is %t and y[%d] is %t\n", i, x.Has(i), i, a.Has(i))
	}
}
