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
}
