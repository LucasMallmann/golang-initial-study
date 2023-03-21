package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No files")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			content, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fmt.Print(content)
			countLines(content, counts)
		}
	}

	for word, counter := range counts {
		fmt.Println("Word:", word, "=>", "Counter:", counter)
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		counts[scanner.Text()]++
	}
}
