package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	arr := make([]int, 512)
	fmt.Println(len(arr), cap(arr))
	arr = append(arr, 1)
	fmt.Println(len(arr), cap(arr))
}
