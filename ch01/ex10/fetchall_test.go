package main

import (
	"fmt"
	"os"
	"testing"
)

var err error

func init() {
	writer, err = os.Create("result.txt")
	if err != nil {
		fmt.Errorf("err! %s", err)
	}
}
func TestFetch(t *testing.T) {
	os.Args = []string{"cmd", "http://gopl.io", "https://www.youtube.com"}
	main()
	main()
}
