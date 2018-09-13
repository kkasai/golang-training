package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type attribute struct {
	name  string
	value string
}

type element struct {
	name       string
	attributes []attribute
}

func parseArgs() []*element {
	var result []*element

	var elm *element
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") {
			if elm == nil {
				fmt.Printf("No element name is specified: [%s] ignored\n", arg)
				continue
			}
			nameValue := strings.Split(arg, "=")
			if len(nameValue) != 2 {
				fmt.Printf("Illegal format: [%s] ignored\n", arg)
			}
			attr := attribute{nameValue[0], nameValue[1]}
			elm.attributes = append(elm.attributes, attr)
			continue
		}

		if elm != nil {
			result = append(result, elm)
		}
		elm = &element{name: arg}
	}
	result = append(result, elm)
	return result
}

func main() {
	element := parseArgs()
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, element) {
				fmt.Printf("%s: %s\n", stackFlattenString(stack), tok)
			}
		}
	}
}

func containsAll(x []xml.StartElement, y []*element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name.Local == y[0].name {
			if containsAllAttributes(x[0].Attr, y[0].attributes) {
				y = y[1:]
			}
		}
		x = x[1:]
	}
	return false
}

func containsAllAttributes(stack []xml.Attr, argAttrs []attribute) bool {
	for _, argAttr := range argAttrs {
		matched := false
		for _, stackAttr := range stack {
			if stackAttr.Name.Local == argAttr.name {
				if stackAttr.Value != argAttr.value {
					return false
				}
				matched = true
			}
			if !matched {
				return false
			}
		}
	}
	return true
}

func stackFlattenString(elements []xml.StartElement) string {
	var result []string
	for _, elem := range elements {
		result = append(result, elem.Name.Local)
	}
	return strings.Join(result, " ")
}