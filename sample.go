package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	m := make(map[string]int, 0)
	m["hello"] = 123
	fmt.Println(m)
	m["hello"]++
	fmt.Println(m)

	if v, ok := m["hello"]; ok {
		fmt.Println(ok, v)
	}

	if v, ok := m["goodbye"]; ok {
		fmt.Println(ok, v)
	}

	mSet := map[int]bool{}
	values := []int{1, 3, 5, 3, 7, 13, 5, 11, 7}
	for _, v := range values {
		mSet[v] = true
	}

	fmt.Printf("%v\n", mSet)
}
