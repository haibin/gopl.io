// Write a function to print the contents of all text nodes in an HTML document tree.
// Do not descend into <script> or <style> elements,
// since their contents are not visible in a web browser.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(doc)
}

func outline(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		return
	}
	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		fmt.Println(strings.TrimSpace(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(c)
	}
}
