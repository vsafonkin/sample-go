package main

import (
	"fmt"
)

type Tester interface{}

func main() {
	fmt.Println("hello")

	var t Tester = 123

	switch t := t.(type) {
	case string:
		fmt.Println("string type, conc2:", t+t)
	case int:
		fmt.Println("int type, mul3:", t*3)
	}
}
