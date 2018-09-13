package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("<%s", e.Type.Local))
	for _, attr := range e.Attr {
		buf.WriteString(fmt.Sprintf(" %s=\"%s\"", attr.Name, attr.Value))
	}
	if len(e.Children) == 0 {
		buf.WriteString("/>\n")
		return buf.String()
	}
	buf.WriteString(">\n")
	for _, child := range e.Children {
		switch n := child.(type) {
		case *Element:
			buf.WriteString(n.String())
		case CharData:
			buf.WriteString(string(n))
		}
	}
	buf.WriteString(fmt.Sprintf("\n</%s>\n", e.Type.Local))
	return buf.String()

}

func main() {
	e, err := build(os.Stdin)
	if err != nil || e == nil {
		fmt.Fprintf(os.Stderr, "xmltree: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", e)
}

func build(r io.Reader) (*Element, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, nil}
			if len(stack) > 0 {
				stack[len(stack) - 1].Children = append(stack[len(stack) - 1].Children, e)
			}
			stack = append(stack, e) // push
		case xml.EndElement:
			if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			stack[len(stack) - 1].Children = append(stack[len(stack) - 1].Children, CharData(tok))
		}
	}
	return nil, nil
}