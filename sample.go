package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	slc := []int{1, 2, 3}
	out := test(slc)

	fmt.Printf("slc: %v\nout: %v\n---\n", slc, out)
	slc[0] = 123
	fmt.Printf("slc: %v\nout: %v\n---\n", slc, out)

	arr := [3]int{7, 8, 9}
	arrOut := getArr(arr)
	fmt.Printf("arr: %v\narrOut: %v\n---\n", arr, arrOut)
	arr[0] = 456
	fmt.Printf("arr: %v\narrOut: %v\n---\n", arr, arrOut)
}

func test(slc []int) []int {
	return slc
}

func getArr(arr [3]int) [3]int {
	return arr
}
