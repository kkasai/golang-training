package main

import (
	"fmt"
	"sort"
)

type PalindromeString []byte

func (x PalindromeString) Len() int           { return len(x) }
func (x PalindromeString) Less(i, j int) bool { return x[i] < x[j] }
func (x PalindromeString) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, max-i, s) {
			return false
		}
	}
	return true
}

func main() {
	result := IsPalindrome(PalindromeString("sator arepo tenet opera rotas"))
	fmt.Println(result)
}
