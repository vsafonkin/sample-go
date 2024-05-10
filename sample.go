package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	Message string
}

func main() {
	fmt.Println("---")

	var err error
	err = CustomErr{
		Message: "hello, alice",
	}
	err = fmt.Errorf("wrap error: %w", err)
	err = fmt.Errorf("second wrap error: %w", err)
	fmt.Println(err)

	fmt.Println(errors.Is(err, CustomErr{}))
	fmt.Println(errors.As(err, &CustomErr{}))
}

func (e CustomErr) Error() string {
	return "custom error"
}
