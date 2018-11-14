package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

func (s *MapIntSet) Has(x int) bool {
	return s.m[x]
}

func (s *MapIntSet) Add(x int) {
	s.m[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for _, x := range t.Ints() {
		s.m[x] = true
	}
}

func (s *MapIntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.m[x] = true
	}
}

func (s *MapIntSet) String() string {
	b := &bytes.Buffer{}
	b.WriteByte('{')
	for i, x := range s.Ints() {
		if i != 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(b, "%d", x)
	}
	b.WriteByte('}')
	return b.String()
}

func (s *MapIntSet) Ints() []int {
	ints := make([]int, 0, len(s.m))
	for x := range s.m {
		ints = append(ints, x)
	}
	sort.IntSlice(ints).Sort()
	return ints
}
