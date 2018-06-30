package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Printf("%d", PopCountFor(uint64(200)))
}

func PopCountFor(x uint64) int {
	var num byte
	for i := 0; i < 8; i++ {
		num += pc[byte(x>>(uint(i)*8))]
	}
	return int(num)
}

func PopCountFormula(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
