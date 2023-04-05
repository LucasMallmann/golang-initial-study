package main

import (
	"fmt"
	"log"
)

func doSomething(nums []int) (*[]int, error) {
	multi := []int{}
	for _, val := range nums {
		multi = append(multi, val*val)
	}
	return &multi, nil
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	result, err := doSomething(values)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*result)
}
