package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error")
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	printElement(images)
	printElement(headings)
}

func printElement(nodes []*html.Node) {
	for _, n := range nodes {
		var attrs string
		for _, a := range n.Attr {
			attrs += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf("<%s%s>\n", n.Data, attrs)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var element []*html.Node
	var f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if n.Data == tag {
					element = append(element, n)
				}
			}
		}
	}
	forEachNode(doc, f)
	return element
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}
