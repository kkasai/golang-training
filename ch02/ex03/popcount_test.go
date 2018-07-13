package main

import "testing"

func TestPopCountFor(t *testing.T) {
	c := PopCountFor(uint64(200))
	if c != 3 {
		t.Errorf("expected: 3 but was actual: %d", c)
	}
}

// 関数が変数をreturnした値を使ってないときや、関数に渡す値が変数でない場合
// コンパイル時に先に計算されてしまう？
var result int
var value = uint64(100)

func BenchmarkPopCountFormula(b *testing.B) {
	var temp int
	for i := 0; i < b.N; i++ {
		temp += PopCountFormula(value)
	}
	result = temp
}

func BenchmarkPopCountFor(b *testing.B) {
	var temp int
	for i := 0; i < b.N; i++ {
		temp += PopCountFor(value)
	}
	result = temp
}
