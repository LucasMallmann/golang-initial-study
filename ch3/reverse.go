package main

import "fmt"

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reversePointer(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j+1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]

	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	newArr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	reverse(arr)
	reversePointer(&newArr)
	fmt.Println(arr)
}
