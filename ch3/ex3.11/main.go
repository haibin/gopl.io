// Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	// Handle floating point
    t := s
	i := strings.LastIndex(s, ".")
	if i != -1 {
		t = s[:i]
	}

    // Handle sign
    sign := ""
    if string(s[0]) == "-" {
        t = t[1:]
        sign = "-"
    }

    r := commaInteger(t)

    return sign + r + s[i:]
}

func commaInteger(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaInteger(s[:n-3]) + "," + s[n-3:]
}