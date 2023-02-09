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

	fmt.Println("--- arrays ---")

	var arr = [3]string{"golang", "python", "javascript"}
	fmt.Println(arr)
	fmt.Printf("Type: %T\n", arr)

	var sl = arr[0:2]
	fmt.Println(sl)
	fmt.Printf("Length: %d\n", len(sl))
	fmt.Printf("Type: %T\n", sl)

	test := []string{"hello", "goodbye"}
	fmt.Println(test)

	// var x []string
	x := []string{"sorry"}
	fmt.Println(x)
}
