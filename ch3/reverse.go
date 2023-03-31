package main

import "fmt"

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func mean(arr []int) float64 {
	summed := sum(arr)
	return float64(summed) / float64(len(arr))
}

func rev(v []int) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 9, 10}

	fmt.Println("Sum is ", sum(arr[:]))
	fmt.Printf("Mean is %.2f", mean(arr[:]))

}
