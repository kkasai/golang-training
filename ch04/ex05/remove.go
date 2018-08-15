package main

import "fmt"

func remove(s []string) []string {
	if len(s) == 0 {
		return s
	}

	current := 0
	for i := 0; i < len(s)-1; i++ {
		if s[current] != s[i+1] {
			s[current+1] = s[i+1]
			current++
		}
	}
	return s[:current+1]
}

func main() {
	s := remove([]string{"aaaa", "bbb", "bbb", "ccc"})
	fmt.Print(s)
}
