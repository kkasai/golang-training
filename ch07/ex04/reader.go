package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type reader struct {
	s string
}

func main() {
	parse("<html><body><a href=\"https://golang.org\" /></body></html>")
}

func parse(s string) {
	doc, err := html.Parse(newReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func newReader(s string) io.Reader {
	return &reader{s}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}
