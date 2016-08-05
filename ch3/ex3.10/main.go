// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if i != len(s)-1 && i%3 == 0 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
