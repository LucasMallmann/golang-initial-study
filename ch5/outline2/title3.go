package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func recoverInPanic() (e error) {
	type bailout struct{}
	var err error
	switch r := recover(); r {
	case nil:
		// no panic
	case bailout{}:
		err = fmt.Errorf("Multiple title elements")
	default:
		panic(r)
	}
	return err
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	})

	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {

}

func forEachNode(node *html.Node, f func(n *html.Node)) {
	if f != nil {
		f(node)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		forEachNode(n, f)
	}
}
