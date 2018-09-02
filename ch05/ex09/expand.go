package main

import (
	"fmt"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(\$\w*)`)

func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func main() {
	fmt.Println(expand("$FoO bar $FUGA", strings.ToLower))
}
