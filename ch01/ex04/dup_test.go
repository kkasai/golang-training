package main

import (
	"bytes"
	"os"
	"testing"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestFileDup(t *testing.T) {
	os.Args = []string{"cmd", "test_data1.txt", "test_data2.txt", "test_data3.txt"}
	main()

	expectedString := "2\taaaaa\ttest_data1.txt test_data3.txt \n2\tccccc\ttest_data1.txt test_data2.txt \n"
	if buffer.String() != expectedString {
		t.Errorf("expected: %s but was actual: %s", expectedString, buffer.String())
	}
}
