package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	w1 := "ぶたがすわらん"
	w2 := "すがわらぶんた"
	fmt.Print(anagram(w1, w2))
}

func anagram(w1, w2 string) bool {
	w1 = strings.ToUpper(w1)
	w2 = strings.ToUpper(w2)
	w1 = strings.Replace(w1, " ", "", -1)
	w2 = strings.Replace(w2, " ", "", -1)
	w1 = SortString(w1)
	w2 = SortString(w2)

	return w1 == w2
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
