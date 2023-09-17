package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	m := make(map[string]int)
	m["bob"] = 123
	m["alisa"] = 456
	m["mike"] = 789
	fmt.Printf("len: %d, map: %v\n", len(m), m)
	fmt.Println("------")
	for k, v := range m {
		fmt.Printf("%s, %d\n", k, v)
	}
}
