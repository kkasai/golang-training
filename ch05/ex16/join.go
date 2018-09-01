package main

import (
	"fmt"
)

func main() {
	strs := []string{"a", "b"}
	str, err := join(",", strs...)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(str)
}

func join(sep string, strs ...string) (string, error) {
	if len(strs) == 0 {
		return "", fmt.Errorf("1つ以上の文字列を渡してください。")
	}
	str := strs[0]
	for _, s := range strs[1:] {
		str += sep + s
	}
	return str, nil
}
