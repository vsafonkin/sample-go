package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	n := 123_456
	x := 0xEA
	b := 0b11101010
	f := 1.23e2
	r := '\uEA12'

	fmt.Println(n, x, b, f, r)

	str := `
	Hello
	"Goodbye"
	\n\t\r
	`

	fmt.Println(str)
}
