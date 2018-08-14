package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(pc[b])
	}
	return total
}

func sha(c1, c2 [sha256.Size]byte) int {
	b := make([]byte, 0, sha256.Size)
	for i := 0; i < sha256.Size; i++ {
		b = append(b, c1[i]^c2[i])
	}
	return popCount(b)
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Print(sha(c1, c2))
}
