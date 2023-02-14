package main

import (
	"fmt"
)

func main() {

	fmt.Println("--- hello ---")
	bnum := 0b11111110
	fmt.Println(bnum)

	xnum := 0xFE
	fmt.Println(xnum)

	unicode := '\u0061'
	fmt.Println(unicode)
	fmt.Println("--- types ---")

	var f float64 = 254.1
	result := int(f)
	fmt.Println(result)
	fmt.Printf("Type: %T\nBinary: %b\nHex: %X\n", result, result, result)
	fmt.Printf("%d\n", 0b11111011)

	users := map[string]int{}
	fmt.Println(users, len(users))
	users["Mike"] = 123
	users["John"] = 456

	fmt.Println(users, len(users))

	var test map[string]string
	if test == nil {
		fmt.Println("test nil")
	}
	fmt.Println(test, len(test))
}
