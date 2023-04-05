package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

func doSomething() *Person {
	person := Person{Name: "Lucas Mallmann", Age: 10}
	return &person
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func doArray(arr []int) {
	arr[0] = 10
}

func main() {
	val := doSomething()
	fmt.Println(val.Age)
	x := []int{0, 1, 2, 3}
	doArray(x)
	fmt.Println(x)
}
