package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(node *html.Node, tags ...string) []*html.Node {
	var nodes []*html.Node
	if len(tags) == 0 {
		return nodes
	}

	var checkNode func(*html.Node)
	checkNode = func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}
		for _, tag := range tags {
			if n.Data == tag {
				nodes = append(nodes, n)
				break
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			checkNode(c)
		}
	}
	checkNode(node)
	return nodes
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Errorf("error whiler parsing html %v", err)
		os.Exit(1)
	}
	// var _images = ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, h := range headings {
		fmt.Printf("<%s>%s<%s>", h.Data, h.FirstChild.Data, h.Data)
	}
}
