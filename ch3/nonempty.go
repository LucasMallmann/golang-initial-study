package main

import "fmt"

func nonempty(s []string) []string {
	i := 0
	for _, val := range s {
		if s[i] != "" {
			s[i] = val
			i++
		}
	}
	return s[:i]
}

func main() {
	values := []string{"One", "Two", "", "Four"}
	nonempty(values)
	fmt.Printf("%q\n", nonempty(values))
	fmt.Printf("%q\n", values)
}
