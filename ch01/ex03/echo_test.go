package ex03

import (
	"os"
	"testing"
	"strings"
)

func BenchmarkIsFor(b *testing.B) {
	os.Args = []string{"cmd", "a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		var s, sep string
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		//fmt.Println(s)
	}
}

func BenchmarkIsJoin(b *testing.B) {
	os.Args = []string{"cmd", "a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		//fmt.Println(strings.Join(os.Args[1:], " "))
		strings.Join(os.Args[1:], " ")
	}
}
