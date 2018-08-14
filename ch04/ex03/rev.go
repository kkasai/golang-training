package main

import "fmt"

const Size = 10

func reverse(s *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [Size]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&s)
	fmt.Print(s)
}
