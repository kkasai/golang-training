package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func listFiles(path string) []string {
	fmt.Println(path)
	var paths []string
	file, err := os.Stat(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	if !file.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print(err)
		return nil
	}
	for _, f := range files {
		paths = append(paths, filepath.Join(path, f.Name()))
	}

	return paths
}

func main() {
	breadthFirst(listFiles, []string{"../"})
}
