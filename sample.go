package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var p *[]int
	fmt.Println(p == nil)
}
