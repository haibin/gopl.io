// Write a function to populate a mapping from element names—p, div,span, and so on—
// to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
    m := make(map[string]int)
	outline(m, doc)
    fmt.Println(m)
}

func outline(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
        m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(m, c)
	}
}
