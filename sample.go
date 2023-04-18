package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	slc := []int{1, 3, 5, 7}
	fmt.Println(sumIntSlice(slc...))
	fmt.Println(sumIntSlice(2, 4, 6))
}

func sumIntSlice(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
