package main

import (
	"fmt"
	"github.com/golang-training/ch07/ex14/eval"
)

func main() {
	expr, err := eval.Parse("min[x, y, z]")
	if err == nil {
		err = expr.Check(map[eval.Var]bool{})
	}
	if err != nil {
		fmt.Println(err)
	}
	env := eval.Env{"x": -10, "y": 4, "z": 10}
	got := fmt.Sprintf("%.6g", expr.Eval(env))
	fmt.Printf("%s\n", expr)
	fmt.Printf("\t%v => %s\n", env, got)
}
