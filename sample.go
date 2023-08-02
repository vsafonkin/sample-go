package main

import (
	"fmt"
)

type TestError struct{}

func main() {
	fmt.Println("hello")

	err := test()
	fmt.Println(err == nil)
}

func test() error {
	var testError TestError
	return testError
}

func (te TestError) Error() string {
	return "test error m"
}
