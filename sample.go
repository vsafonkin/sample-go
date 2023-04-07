package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	r := '\u212E'
	fmt.Printf("%c %[1]b\n", r)
	fmt.Printf("%c %[1]b\n", r<<1)

}

func DoSomething() {
	fmt.Println("Do something")
}
