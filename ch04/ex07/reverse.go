package main

import "fmt"

func reverse(bytes []byte) []byte {
	runes := []rune(string(bytes))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	copy(bytes, []byte(string(runes)))

	return bytes
}

func main() {
	fmt.Printf("%s", reverse([]byte("abcdef")))
}
