package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	scanner   *bufio.Scanner
	algorythm *string
)

func init() {
	algorythm = flag.String("a", "sha256", "supports sha256, sha384 and sha512.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func main() {

	printFuncs := map[string]func(string){
		"sha256": printSHA256,
		"sha384": printSHA384,
		"sha512": printSHA512,
	}

	flag.Parse()
	f, ok := printFuncs[*algorythm]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	for {
		if input, eof := eGetLine(); eof {
			break
		} else {
			f(input)
		}
	}
}

func printSHA256(input string) {
	fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
}
func printSHA384(input string) {
	fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
}
func printSHA512(input string) {
	fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
}

func eGetLine() (string, bool) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}
		return "", true
	}
	return scanner.Text(), false
}
