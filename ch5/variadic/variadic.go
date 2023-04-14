package main

import "fmt"

func min(values ...int) int {
	println(values)
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func join(values ...string) string {
	value := ""
	for _, v := range values {
		value += fmt.Sprintf("%s ", v)
	}
	return value
}

func main() {
	println(min(9, 2912, 92389, 192, -1))
	println(join("oi", "tudo", "bem?"))
}
