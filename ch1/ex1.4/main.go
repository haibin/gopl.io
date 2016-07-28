// Modify dup2 to print the names of all files in whicheach duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counts []string
	files := os.Args[1:]
	if len(files) == 0 {
		if hasDuplicateLines(os.Stdin) {
			counts = append(counts, "Stdin")
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			if hasDuplicateLines(f) {
				counts = append(counts, arg)
			}
			f.Close()
		}
	}
	for _, file := range counts {
		fmt.Printf("%s\n", file)
	}
}

func hasDuplicateLines(f *os.File) bool {
	counts := make(map[string]int)

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()

	for _, n := range counts {
		if n > 1 {
			return true
		}
	}

	return false
}
