package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	b := 0b11100100
	x := 0xE4
	num := 228
	fmt.Println(b, x)
	fmt.Printf("decimal: %d, bin: %#[1]b, hex:%#[1]x\n", num)

	f := 1.234e2
	fmt.Printf("Type: %T, value: %#[1]f\n", f)

	r := '\U000000BE'
	fmt.Println(r)
	fmt.Printf("Type: %T, value: %#[1]q\n", r)
	str := "\u212E\x41\U0000212E"
	fmt.Println(str)
}
