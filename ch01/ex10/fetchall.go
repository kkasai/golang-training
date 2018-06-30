package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(writer, <-ch)
	}

	fmt.Fprintf(writer, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	s := fmt.Sprintf("%s_%s.txt", resp.Request.URL.Host, start.Format("20060102150405.99"))
	file, err := os.Create(s)
	if err != nil {
		fmt.Sprintf("failure create file: %v", err)
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s output_filename: %s", secs, nbytes, url, s)
}
