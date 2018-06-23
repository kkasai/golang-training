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
	for i := 1; i < len(os.Args); i++ {
		s := fmt.Sprintf("%d %s", i, os.Args[i])
		fmt.Fprintln(writer, s)
	}
}