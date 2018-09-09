package main

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	data := []string{
		"Hello World\n",
		"Go言語\n",
	}

	w, c := CountingWriter(os.Stdout)

	var total int64 = 0

	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("expected: %d but was actual: %d", total, *c)
		}
	}
}
