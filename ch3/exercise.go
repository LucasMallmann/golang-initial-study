package main

import (
	"bytes"
	"fmt"
	"sort"
)

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, val := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteString("]")
	return buf.String()
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}

func anagram(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})

	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes1) == string(runes2)
}

func main() {
	fmt.Println(anagram("ola", "ala"))
}
