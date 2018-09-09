package main

import (
	"fmt"
	"io"
	"os"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper
	wp.w = w
	return &wp, &(wp.c)
}

type Wrapper struct {
	c int64
	w io.Writer
}

func (wp *Wrapper) Write(b []byte) (n int, err error) {
	n, err = wp.w.Write(b)
	wp.c += int64(n)
	return
}

func main() {
	w, c := CountingWriter(os.Stdout)
	w.Write([]byte("hello world\n"))
	fmt.Println("count:", *c)
}
