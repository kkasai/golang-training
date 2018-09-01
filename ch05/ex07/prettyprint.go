package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"io"
	"strings"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startNode, endNode)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	fmt.Fprintf(writer, "%*s<%s%s%s\n", depth*2, "", n.Data, attrStr, end)

	depth++
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Fprintf(writer, "%*s</%s>\n", depth*2, "", n.Data)
}

func textNode(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	fmt.Fprintf(writer, "%*s%s\n", depth*2, "", n.Data)
}

func commentNode(n *html.Node) {
	fmt.Fprintf(writer, "%*s<!--%s-->\n", depth*2, "", n.Data)
}

func startNode(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		textNode(n)
	case html.CommentNode:
		commentNode(n)
	}
}

func endNode(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		endElement(n)
	}
}
