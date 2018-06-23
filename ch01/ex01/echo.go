package main

import (
	"os"
	"fmt"
	"io"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}

func main()  {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Fprintln(writer, s)
}