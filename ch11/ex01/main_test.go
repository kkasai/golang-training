package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	for _, test := range []struct {
		bytes   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("Hi, 世."),
			map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '世': 1, '.': 1},
			[]int{0, 5, 0, 1, 0},
			0,
		}, {
			[]byte("Hi, 世.\300"),
			map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '世': 1, '.': 1},
			[]int{5, 5, 0, 1, 0},
			1,
		},
	} {
		counts, utflen, invalid, err := charCount(bytes.NewReader(test.bytes))
		if err != nil {
			t.Error(err)
			continue
		}

		for k, v := range test.counts {
			count, ok := counts[k]
			if !ok {
				t.Errorf("%c is not included\n", k)
				continue
			}
			if count != v {
				t.Errorf("count for %c is %d, but want %d\n", k, count, v)
				continue
			}
		}

		for i := 1; i < len(utflen); i++ {
			if utflen[i] != test.utflen[i] {
				t.Errorf("utflen[%d] is %d, but want %d\n", i, utflen[i], test.utflen[i])
				continue
			}
		}

		if invalid != test.invalid {
			t.Errorf("invalid is %d, but want %d\n", invalid, test.invalid)
		}
	}
}
