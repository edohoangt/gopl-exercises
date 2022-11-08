package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)


func comma(s string) string {
	var buf bytes.Buffer

	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.LastIndex(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3

	if pre == 0 {
		pre = 3
	}
	buf.WriteString(mantissa[:pre])

	for i, c := range mantissa[pre:] {
		if i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}

	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}