package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)

	bnum := 0b11111110
	fmt.Println(bnum)

	xnum := 0xFE
	fmt.Println(xnum)

	unicode := '\u0061'
	fmt.Println(unicode)

	var f float64 = 254.1
	result := int(f)
	fmt.Println(result)
	fmt.Printf("Type: %T\nBinary: %b\nHex: %X\n", result, result, result)
	fmt.Printf("%d\n", 0b11111011)

	myFunc := test
	myFunc("hello", "goodbye")

	err := runSomeFunc(myFunc, "abrakadabra")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var str string

	id := 123
	str = fmt.Sprint(id)
	fmt.Println(str)
}

func test(strarray ...string) error {
	for _, value := range strarray {
		fmt.Println(value)
	}
	return nil
}

func runSomeFunc(someFunc func(...string) error, str string) error {
	err := someFunc(str)
	return err
}
