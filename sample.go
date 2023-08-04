package main

import (
	"fmt"
	"io"
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
	return 123, "hello", io.EOF
}
