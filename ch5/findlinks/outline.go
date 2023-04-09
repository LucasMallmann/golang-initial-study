package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tree elements: %v", err)
		os.Exit(1)
	}
	elementsTree := tree(nil, doc)
	for _, el := range elementsTree {
		fmt.Println(el)
	}
}

func tree(elements []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		elements = append(elements, node.Data)
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		elements = tree(elements, n)
	}
	return elements
}
