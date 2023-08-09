package main

import (
	"fmt"
)

type TestError struct {
	Test      string
	Message   string
	ErrStatus int
}

func (e TestError) Error() string {
	return fmt.Sprintf("error message: %s, error status: %d\n", e.Message, e.ErrStatus)
}

func main() {
	fmt.Println("-------------")
}
