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

	// test := math.Pow(2, 64)
	var test float64 = 254.1
	result := int(test)
	fmt.Println(result)
	fmt.Printf("Type: %T\nBinary: %b\nHex: %x\n", result, result, result)
}
