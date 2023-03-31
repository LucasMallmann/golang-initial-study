package main

import "fmt"

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fibo(n int) int {
	if n == 1 {
		return n
	}
	return fibo(n-1) * n
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	var gcd int = gcd(10, 2)
	fmt.Println(gcd)
	fmt.Println(fibo(5))
}
