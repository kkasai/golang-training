package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text_node: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Data == "script" || n.Data == "style" {
		return links
	}

	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			links = append(links, s)
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
