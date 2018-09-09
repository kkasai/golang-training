package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(strings.NewReader(string(p)))
	input.Split(bufio.ScanWords)
	var count int
	for input.Scan() {
		count++
	}
	*c += WordCounter(count)
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(strings.NewReader(string(p)))
	var count int
	input.Split(bufio.ScanLines)
	for input.Scan() {
		count++
	}
	*c += LineCounter(count)
	return len(p), nil
}
func main() {
	var wc WordCounter
	wc.Write([]byte("hello world golang"))
	fmt.Println("word: ", wc)

	wc = 0
	var wordName = "Dolly"
	fmt.Fprintf(&wc, "hello, %s", wordName)
	fmt.Println("word: ", wc)

	var rc LineCounter
	rc.Write([]byte("hello\nworld\ngolang"))
	fmt.Println("row: ", rc)

	rc = 0
	var rowName = "Dolly"
	fmt.Fprintf(&rc, "hello,\n%s", rowName)
	fmt.Println("row: ", rc)
}
