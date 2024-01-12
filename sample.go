package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	var arr [3]int
	fmt.Println(arr)

	arr[1] = 123
	fmt.Println(arr)

	func(arr [3]int) {
		arr[2] = 456
	}(arr)

	fmt.Println(arr)
}
