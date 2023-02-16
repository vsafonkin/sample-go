package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("--- hello ---")

	var currentTime = time.Now()
	fmt.Println(currentTime)

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

	type user struct {
		name string
		age  int
	}

	var mike user
	mike.name = "mike"
	mike.age = 23

	bob := user{
		name: "bob",
		age:  34,
	}

	fmt.Println(mike, bob)

	order := struct {
		id    int
		title string
	}{
		id:    12345,
		title: "pc",
	}

	fmt.Println(order)
}
