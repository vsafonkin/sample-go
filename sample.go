package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")
	slc := []int{1, 2, 3}
	setFirst(slc)
	fmt.Println(slc)

	m := map[string]int{
		"hello":   1,
		"goodbye": 2,
	}

	mHelloKey(m)
	fmt.Printf("%+v\n", m)
}

func setFirst(slc []int) error {
	if len(slc) == 0 {
		return fmt.Errorf("slice len is zero")
	}
	slc[0] = 123
	return nil
}

func mHelloKey(m map[string]int) error {
	if len(m) == 0 {
		return fmt.Errorf("empty map")
	}
	m["hello"] = 789
	return nil
}
