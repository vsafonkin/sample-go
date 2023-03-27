package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	slc := []int{7, 8, 9}
	err := setFirstElement(slc, 123)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slc)
}

func setFirstElement(slice []int, num int) error {
	if slice == nil {
		return fmt.Errorf("error! slice is nil")
	}
	if len(slice) == 0 {
		return fmt.Errorf("error! slice len is zero")
	}
	slice[0] = num
	return nil
}
