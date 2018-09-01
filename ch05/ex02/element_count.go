package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "element count: %v\n", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	for e, n := range visit(counts, doc) {
		fmt.Printf("%q\t%d\n", e, n)
	}
}

func visit(counts map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counts
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	counts = visit(counts, n.FirstChild)
	counts = visit(counts, n.NextSibling)

	return counts
}
