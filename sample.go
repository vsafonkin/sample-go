package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	var arr []int = []int{0, 0, 0}
	fmt.Println(arr)

	arr[1] = 123
	fmt.Println(arr)

	func(arr []int) {
		arr[2] = 456
		fmt.Println(cap(arr))
		arr = append(arr, 789)
		fmt.Println(cap(arr))
	}(arr)

	fmt.Println(arr)
}
