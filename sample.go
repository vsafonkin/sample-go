package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var r rune
	fmt.Printf("%d %[1]c %[1]T\n", r)
}
