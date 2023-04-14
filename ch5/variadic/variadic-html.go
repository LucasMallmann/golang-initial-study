package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var count = 0
var inter = 0

func contains(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}

func elementsByTagName(node *html.Node, tags ...string) []*html.Node {
	foundNodes := []*html.Node{}
	var searchForTags func(n *html.Node)
	searchForTags = func(n *html.Node) {
		if n == nil {
			return
		}
		tagName := n.Data
		if contains(tags, tagName) {
			foundNodes = append(foundNodes, n)
		}
	}
	forEachNode(node, searchForTags)
	return foundNodes
}

func forEachNode(node *html.Node, f func(n *html.Node)) {
	if f != nil {
		f(node)
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		forEachNode(n, f)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	foundNodes := elementsByTagName(doc, "a")
	for _, node := range foundNodes {
		tag, content := node.Data, node.FirstChild.Data
		fmt.Printf("<%s>%s<%s>\n", tag, content, tag)
	}
}
