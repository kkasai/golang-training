package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(comma("111121"))
}

func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}

	var buf bytes.Buffer
	i := n % 3
	if i == 0 {
		i = 3
	}
	for _, v := range []byte(str) {
		if i == 0 {
			buf.WriteString(",")
			i = 3
		}
		buf.WriteByte(v)
		i--
	}
	return buf.String()
}
