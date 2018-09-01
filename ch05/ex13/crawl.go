package main

import (
	"log"
	"os"

	"fmt"
	"gopl.io/ch5/links"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var hostName string

func download(urlString string) error {
	url, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if hostName == "" {
		hostName = url.Host
	}
	if hostName != url.Host {
		return nil
	}
	dir := url.Host
	var filename string
	if filepath.Ext(filename) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(urlString)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func crawl(url string) []string {
	err := download(url)
	if err != nil {
		fmt.Println(err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: crawl <url> [<url>...]")
		os.Exit(1)
	}
	
	breadthFirst(crawl, os.Args[1:])
}
