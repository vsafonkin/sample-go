package main

import (
	"fmt"
)

type User struct{}

func main() {
	fmt.Println("hello")
	doSomething()

	ch := make(chan int)
	fmt.Println(ch)
}

func doSomething() {
	fmt.Println("do something")
}
