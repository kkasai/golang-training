package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
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

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			links = addLink(links, n, "href")
		case "img", "script":
			links = addLink(links, n, "src")
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}

func addLink(links []string, n *html.Node, key string) []string {
	for _, a := range n.Attr {
		if a.Key == key {
			links = append(links, a.Val)
		}
	}
	return links
}
