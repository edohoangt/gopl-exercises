package main

import (
	"bytes"
	"fmt"
	"os"
)


func comma(s string) string {
	start := len(s) % 3
	var buf bytes.Buffer

	if start == 0 {
		start = 3
	}
	buf.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[start:start+3])
	}

	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}