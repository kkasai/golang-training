package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	in := bufio.NewReader(os.Stdin)
	invalid := 0
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsControl(r):
			counts["control"]++
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsPunct(r):
			counts["punctuation"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsSymbol(r):
			counts["symbol"]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for s, n := range counts {
		fmt.Printf("%s\t%d\n", s, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
