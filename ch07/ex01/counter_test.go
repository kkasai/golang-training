package main

import "testing"

func TestWordCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 2},
		{"Hello World golang", 3},
		{"こんにちは　世界　Go　言語", 4},
	}

	var c WordCounter
	for _, d := range data {
		c = 0

		bytes := []byte(d.s)
		n, err := c.Write(bytes)

		if err != nil {
			t.Errorf("Unpexected Error : %v", err)
			continue
		}

		if n != len(bytes) {
			t.Errorf("expected: %d but was actual: %d", len(bytes), n)
			continue
		}

		if c != WordCounter(d.expected) {
			t.Errorf("expected: %d but was actual: %d", d.expected, c)
		}
	}
}

func TestLineCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 1},
		{"Hello World\nHello World", 2},
		{"Hello World\nこんにちは\n世界", 3},
	}

	var c LineCounter
	for _, d := range data {
		c = 0

		bytes := []byte(d.s)
		n, err := c.Write(bytes)

		if err != nil {
			t.Errorf("Unpexected Error : %v", err)
			continue
		}

		if n != len(bytes) {
			t.Errorf("expected: %d but was actual: %d", len(bytes), n)
			continue
		}

		if c != LineCounter(d.expected) {
			t.Errorf("expected: %d but was actual: %d", d.expected, c)
		}
	}
}
