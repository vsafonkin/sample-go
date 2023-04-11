package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")
	slc := []int{1, 2, 3}
	setFirst(slc)
	fmt.Println(slc)
}

func setFirst(slc []int) error {
	if len(slc) == 0 {
		return fmt.Errorf("slice len is zero")
	}
	slc[0] = 123
	return nil
}
