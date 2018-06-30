package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	var lines []string
	for line := range counts {
		lines = append(lines, line)
	}
	sort.Strings(lines)
	for _, line := range lines {
		var filenameArray []string
		for tmpFilename := range counts[line] {
			filenameArray = append(filenameArray, tmpFilename)
		}
		sort.Strings(filenameArray)
		var sum int
		var filenames string
		for _, filename := range filenameArray {
			sum += counts[line][filename]
			filenames += filename + " "
		}
		if sum > 1 {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", sum, line, filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}
