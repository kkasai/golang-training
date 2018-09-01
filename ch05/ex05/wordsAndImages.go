package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordsAndImages: %v\n", err)
			continue
		}

		fmt.Printf("[%s]\n=> Words : %d, Images : %d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}

	input := bufio.NewScanner(strings.NewReader(n.Data))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		images += i
		words += w
	}

	return
}
