package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var arr []int

	for i := range [1_000_000]struct{}{} {
		temp := cap(arr)
		arr = append(arr, i)
		if temp != cap(arr) {
			fmt.Printf("index: %d, len: %d, cap: %d\n", i, len(arr), cap(arr))
		}
	}

}
