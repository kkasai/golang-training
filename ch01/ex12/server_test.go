package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go main()

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	res, err := client.Get("http://localhost:8000")
	if err != nil {
		fmt.Errorf("error! %v", err)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "image/gif" {
		t.Errorf("expected: image/gif but was actual: %s", contentType)
	}

	res2, err2 := client.Get("http://localhost:8000?cycles=20")
	if err2 != nil {
		fmt.Errorf("error! %v", err)
	}

	contentType = res2.Header.Get("Content-Type")
	if contentType != "image/gif" {
		t.Errorf("expected: image/gif but was actual: %s", contentType)
	}
}
