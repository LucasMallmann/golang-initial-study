package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v", err)
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				counts[line]++
			}
		}
	}

	for line, counter := range counts {
		if counter > 1 {
			fmt.Printf("%d\t%s", counter, line)
		}
	}
}
