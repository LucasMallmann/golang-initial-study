package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func forEachNode(node *html.Node, pre func(n *html.Node), post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = node.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

var level int

func preElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", level*2, "", node.Data)
		level++
	}
}

func postElement(node *html.Node) {
	if node.Type == html.ElementNode {
		level--
		fmt.Printf("%*s</%s>\n", level*2, "", node.Data)
	}
}

func outline(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(doc.Data)
	forEachNode(doc, preElement, postElement)
}
