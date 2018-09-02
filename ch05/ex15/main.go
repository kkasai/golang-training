package main

import (
	"fmt"
	"os"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("スライスの長さが0です。")
	}

	m := vals[0]
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m, nil
}

func max2(val int, vals ...int) int {
	m := val
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("スライスの長さが0です。")
	}

	m := vals[0]
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m, nil
}

func min2(val int, vals ...int) int {
	m := val
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m
}

func main() {
	values := []int{1, 2, 3, 4}
	max, err := max(values...)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println("max(): ", max)
	fmt.Println("max2(): ", max2(5, values...))

	min, err := min(values...)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println("min(): ", min)
	fmt.Println("min2(): ", min2(5, values...))
}
