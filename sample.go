package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sample")

	bnum := 0b11111110
	fmt.Println(bnum)

	xnum := 0xFE
	fmt.Println(xnum)

	unicode := '\u0061'
	fmt.Println(unicode)

	var bytenum byte
	fmt.Println(bytenum)

	fmt.Println(10 / 3)
	fmt.Println(10 % 3)

	b := 0b101
	fmt.Println(b)

	fmt.Println(bnum + xnum)
	fmt.Println("-----")

	var x int = 254
	var y float64 = 12.3
	result := x + int(y)
	fmt.Println("Result:", result)
}
