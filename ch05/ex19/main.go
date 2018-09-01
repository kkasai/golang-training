package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()
	panic(2)
}
