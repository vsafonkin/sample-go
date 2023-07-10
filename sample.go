package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hello")

	n := -1
	if n := rand.Intn(100); n > 0 {
		fmt.Println(n)
	}

	fmt.Println(n)
}
