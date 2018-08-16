package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	count := make(map[string]int)

	for input.Scan() {
		count[input.Text()] += 1
	}

	for s, i := range count {
		fmt.Printf("%s\t%d\n", s, i)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s files\n", os.Args[0])
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		fmt.Println(f)
		wordfreq(f)
	}
}
