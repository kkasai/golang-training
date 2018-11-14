package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountFor(x uint64) int {
	var num int
	for i := 0; i < 64; i++ {
		if x>>uint64(i)&1 == 1 {
			num++
		}
	}
	return num
}

func PopCount64Shift(x uint64) int {
	var num int
	for i := 0; i < 64; i++ {
		if x>>uint64(i)&1 == 1 {
			num++
		}
	}
	return num
}

func PopCountOne(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
