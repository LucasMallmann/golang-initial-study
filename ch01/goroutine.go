package main

import (
	"fmt"
	"time"
)

func sayHi(s string) {
	for i := 1; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ":", i)
	}
}

func main() {
	go sayHi("First Goroutine")
	go sayHi("Second Goroutine")

	time.Sleep(time.Second)

	fmt.Println("Main Goroutine is done")
}
