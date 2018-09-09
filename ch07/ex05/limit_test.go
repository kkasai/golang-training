package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	for _, tc := range []struct {
		data  string
		limit int
	}{
		{"0123456789", 3},
		{"0123456789", 0},
		{"0123456789", 3},
	} {
		lr1 := LimitReader(strings.NewReader(tc.data), tc.limit)
		lr2 := io.LimitReader(strings.NewReader(tc.data), int64(tc.limit))

		b1 := &bytes.Buffer{}
		n1, _ := b1.ReadFrom(lr1)
		b2 := &bytes.Buffer{}
		n2, _ := b2.ReadFrom(lr2)

		if n1 != n2 {
			t.Logf("n=%d", n1)
			t.Fail()
		}
		if b1.String() != b2.String() {
			t.Logf(`"%s" != "%s"`, b1.String(), b2.String())
		}
	}
}
