package main

import "fmt"

func main() {
	x, y := 10, 15

	fmt.Printf("Value before swaping X -> %d\n", x)
	fmt.Printf("Value before swaping Y -> %d\n", y)

	swap(&x, &y)

	fmt.Printf("Value after swap X -> %d\n", x)
	fmt.Printf("Value after swap Y -> %d\n", y)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
