package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type packageInfo struct {
	ImportPath string
	Name string
	Deps []string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ./deps package[s]")
	}
	packages, err := executeGoList(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("dependency\n")
	for _, p := range packages {
		fmt.Printf("%s: %v\n", p.Name, p.Deps)
	}
	allPackages, err := executeGoList("...")
	if err != nil {
		log.Fatal(err)
	}
	depPackage := map[string][]string{}
	for _, p := range allPackages {
		for _, arg := range os.Args[1:] {
			for _, d := range p.Deps {
				if d == arg {
					depPackage[arg] = append(depPackage[arg], p.ImportPath)
				}
			}
		}
	}
	fmt.Print("\n")
	for k, v := range depPackage {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func executeGoList(args ...string) ([]*packageInfo, error) {
	cmdArgs := []string{"list", "-json"}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("go", cmdArgs...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	jsonDecoder := json.NewDecoder(bytes.NewReader(output))
	var pInfos []*packageInfo
	for {
		var pInfo packageInfo
		err := jsonDecoder.Decode(&pInfo)
		if err != nil {
			if err == io.EOF {
				return pInfos, nil
			}
			log.Fatal(err)
		}
		pInfos = append(pInfos, &pInfo)
	}
}