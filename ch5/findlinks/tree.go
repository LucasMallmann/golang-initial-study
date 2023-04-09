package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	els := tree(nil, doc)
	for _, el := range els {
		fmt.Println(el)
	}
}

func tree(elements []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		elements = append(elements, node.Data)
	}

	for n := node.FirstChild; n != nil; n = node.NextSibling {
		elements = tree(elements, n)
	}
	return elements
}
