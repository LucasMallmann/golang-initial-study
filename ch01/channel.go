package main

import "fmt"

func greet(c chan string) {
	name := <-c
	fmt.Println("Name is :", name)
}

func main() {
	c := make(chan string)
	go greet(c)
	c <- "John"
	close(c)
}
