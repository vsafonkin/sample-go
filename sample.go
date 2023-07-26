package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, pointers")

	arr := [3]int{1, 2, 3}
	arrCopy := arr

	fmt.Printf("arr: %v\narrCopy: %v\n", arr, arrCopy)
	arr[0] = 789
	fmt.Printf("arr: %v\narrCopy: %v\n", arr, arrCopy)

	slc := []int{1, 2, 3}
	slcCopy := slc

	fmt.Printf("slc: %v\nslcCopy: %v\n", slc, slcCopy)
	slc[0] = 789
	fmt.Printf("slc: %v\nslcCopy: %v\n", slc, slcCopy)
	slcCopy[0] = 456
	fmt.Printf("slc: %v\nslcCopy: %v\n", slc, slcCopy)
}
