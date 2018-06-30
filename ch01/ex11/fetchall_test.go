package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Minute)
})

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	urls = append(urls, ts.URL)
	os.Args = append(os.Args, urls...)
	main()
}
