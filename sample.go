package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	var slc []int
	addItem(5, slc)
	fmt.Println(slc)
}

func addItem(item int, slc []int) {
	slc = append(slc, item)
	fmt.Println("addItem slc:", slc)
}
