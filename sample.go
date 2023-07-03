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
}
