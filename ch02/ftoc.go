package main

import "fmt"

func main() {
	// Converts Fahrenheit to Celsius
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(value float64) float64 {
	return (value - 32) * 5 / 9
}
