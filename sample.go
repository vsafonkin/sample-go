package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	arr := []int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl = append(sl, 6)
	fmt.Printf("arr: %v, sl: %v\n", arr, sl)
	sl[0] = 7
	fmt.Printf("arr: %v, sl: %v\n", arr, sl)
	arr[1] = 8
	fmt.Printf("arr: %v, sl: %v\n", arr, sl)
}
