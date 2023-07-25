package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, pointers")

	var p *int
	fmt.Println(p == nil)
	fmt.Printf("%p\n", p)
	updatePointer(p)
	fmt.Printf("%p\n", p)
}

func updatePointer(p *int) {
	fmt.Printf("%p, %t\n", p, p == nil)
	temp := 123
	p = &temp
}
