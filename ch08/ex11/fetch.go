package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})
var wg sync.WaitGroup

func cancelled(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch) // start a goroutine
	}
	fmt.Println(<-ch) // receive from channel ch
	close(done)       // cancel others
	wg.Wait()
}

func fetch(url string, ch chan<- string) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if cancelled(done) {
		return
	}

	cancelChan := make(chan struct{})
	req.Cancel = cancelChan

	go func() {
		select {
		case <-done:
			close(cancelChan)
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case ch <- fmt.Sprint(err): // send to channel ch
		case <-done:
		}
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	select {
	case ch <- fmt.Sprintf("%d  %s", nbytes, url):
	case <-done:
	}
}