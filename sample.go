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

	var test float64 = 254.1
	result := int(test)
	fmt.Println(result)
	fmt.Printf("Type: %T\nBinary: %b\nHex: %X\n", result, result, result)
	fmt.Printf("%d\n", 0b11111011)

	fmt.Println("--- arrays ---")

	var arr = [3]string{"golang", "python", "javascript"}
	fmt.Println(arr)
}
