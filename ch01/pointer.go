package main

import "fmt"

var pointer *int

func main() {
	i, j := 42, 43
	p := &i //points to i

	fmt.Println(*p)

	*p = 40
	fmt.Println(i)

	p = &j      // point to j
	*p = *p / 2 // divide j through the pointer
	fmt.Println(j)
}
