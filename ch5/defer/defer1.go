package main

import "fmt"

func f(n int) {
	fmt.Printf("f(%d)\n", n+0/n) // panics if n == 0
	defer fmt.Printf("defer %d\n", n)
	f(n - 1)
}

func printAllOperations(x, y int) {
	type divideByZero struct{}

	defer func() {
		switch r := recover(); r {
		case nil:
			// No error

		case divideByZero{}:
			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
			fmt.Println("Proceeding to alternative flow skipping division.")
			printOperationsSkipDivide(x, y)

		default:
			panic(r)
		}
	}()

	if x == 0 {
		panic(divideByZero{})
	}

	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("sum=%v, divide=%v, multiply=%v \n", sum, divide, multiply)
}

func printOperationsSkipDivide(x int, y int) {
	sum, multiply := x+y, y*x
	fmt.Printf("sum=%v, multiply=%v \n", sum, multiply)
}

func main() {
	x := 0
	y := 20
	printAllOperations(x, y)
}
