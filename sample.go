package main

import (
	"fmt"
)

const (
	One = iota
	Two
	Three
)

func main() {
	fmt.Println("-----")

	fmt.Printf("%d, %d, %d\n", One, Two, Three)
}
