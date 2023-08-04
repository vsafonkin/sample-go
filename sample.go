package main

import (
	"fmt"
)

type TestError struct {
	Message string
}

func (e TestError) Error() string {
	return e.Message
}

func main() {
	fmt.Println("hello")

	err := wrap()
	fmt.Println(err == nil)
}

func doSomething() *TestError {
	return nil
}

func wrap() error {
	return doSomething()
}
