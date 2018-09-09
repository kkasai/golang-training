package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	for _, test := range []struct {
		expr string
		env  Env
		want string
	}{
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"-1 + -x", Env{"x": 3}, "-4"},
		{"-1 * -x", Env{"x": 3}, "3"},
	} {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		fmt.Printf("%s\n", expr.String())
		reparsedExpr, err := Parse(expr.String()) // parse again
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		got := fmt.Sprintf("%.6g", reparsedExpr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
