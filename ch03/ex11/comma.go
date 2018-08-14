package main

import (
	"fmt"
	"strings"
)

func main() {
	str := comma("-1111111.123")
	fmt.Printf(str)
}

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + comma(s[1:])
	}

	pi := strings.IndexByte(s, '.')
	if pi >= 0 {
		return comma(s[:pi]) + s[pi:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
