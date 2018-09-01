package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: findElement url id")
		os.Exit(1)
	}

	findElement(os.Args[1], os.Args[2])
}

func findElement(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := forEachNode(doc, id, startElement, nil)
	var attrs string
	for _, a := range node.Attr {
		attrs += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
	}
	fmt.Printf("<%s%s>", node.Data, attrs)

	return nil
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}

	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if ElementByID(n, id) != nil {
			return true
		}
	}
	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	for _, a := range doc.Attr {
		if a.Key == "id" && a.Val == id {
			return doc
		}
	}
	return nil
}
