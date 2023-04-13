package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	slc := []int{1, 2, 3}
	addItem(5, slc)
	fmt.Println(slc)
}

func addItem(item int, slc []int) {
	slc = append(slc, item)
}
