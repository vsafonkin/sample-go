package main

import (
	"fmt"
)

type Counter interface {
	count() int
}

func main() {
	fmt.Println("hello")
	doSomething()
}

func doSomething() (int, string, error) {
	fmt.Println("Do something")
	return 123, "hello", nil
}
