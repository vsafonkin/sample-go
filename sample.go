package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	slc := make([]int, 1)

	func(slc []int) {
		slc = append(slc, 123)
	}(slc)

	fmt.Println(slc)
}
