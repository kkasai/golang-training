package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang-training/ch02/ex02/distanceconv"
	"github.com/golang-training/ch02/ex02/tempconv"
	"github.com/golang-training/ch02/ex02/weightconv"
	"io"
	"os"
	"strconv"
)

var writer io.Writer

func init() {
	writer = os.Stdout
}

type stringFlag struct {
	set   bool
	value string
}

func (ff *stringFlag) Set(x string) error {
	ff.value = x
	ff.set = true
	return nil
}

func (ff *stringFlag) String() string {
	return ff.value
}

var tStringFlag stringFlag
var dStringFlag stringFlag
var wStringFlag stringFlag

func init() {
	flag.Var(&tStringFlag, "t", "温度")
	flag.Var(&dStringFlag, "d", "長さ")
	flag.Var(&wStringFlag, "w", "重さ")
}

func main() {
	flag.Parse()
	if flag.NFlag() < 1 {
		var kind string
		input := bufio.NewScanner(os.Stdin)
		for kind != "t" && kind != "w" && kind != "d" {
			fmt.Println("種類を選択してください。")
			fmt.Println("[t]:温度 [d]:長さ [w]:重さ")
			if input.Scan() {
				kind = input.Text()
			}
		}
		var value float64
		var err error
		for {
			fmt.Println("数値を入力してください。")
			if input.Scan() {
				value, err = strconv.ParseFloat(input.Text(), 64)
				if err == nil {
					break
				}
			}
		}
		printConv(kind, value)
		os.Exit(0)
	}

	if tStringFlag.set {
		t := parseFloat64(tStringFlag.value)
		printConv("t", t)
	}

	if dStringFlag.set {
		d := parseFloat64(dStringFlag.value)
		printConv("d", d)
	}

	if wStringFlag.set {
		w := parseFloat64(wStringFlag.value)
		printConv("w", w)
	}
}

func printConv(kind string, value float64) {
	switch kind {
	case "t":
		f := tempconv.Fahrenheit(value)
		c := tempconv.Celsius(value)
		fmt.Fprintf(writer, "%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	case "d":
		m := distanceconv.Metre(value)
		f := distanceconv.Feet(value)
		fmt.Fprintf(writer, "%s = %s, %s = %s\n",
			m, distanceconv.MToF(m), f, distanceconv.FToM(f))
	case "w":
		p := weightconv.Pound(value)
		k := weightconv.Kilogram(value)
		fmt.Fprintf(writer, "%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}

func parseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return f
}
