package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-training/ch07/ex16/eval"
)

func main() {
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values, ok := r.Form["expr"]
	if !ok {
		http.Error(w, "no expr", http.StatusBadRequest)
		return
	}

	for _, v := range values {
		expr, err := eval.Parse(v)
		if err != nil {
			http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
			return
		}

		result := expr.Eval(eval.Env{})
		fmt.Fprintf(w, "%s = %g\n", v, result)
	}
}
