package main

import (
	"fmt"
	"math"
	"os"
)
import "github.com/golang-training/ch07/ex13/eval"

func main() {
	expr, err := eval.Parse("sqrt(A / pi)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", expr.String())
	reparsedExpr, err := eval.Parse(expr.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	env := eval.Env{"A": 87616, "pi": math.Pi}
	got := fmt.Sprintf("%.6g", reparsedExpr.Eval(env))
	fmt.Printf("\t%v => %s\n", env, got)
}
