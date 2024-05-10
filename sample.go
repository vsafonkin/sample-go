package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	Message string
}

func main() {

	var err error
	err = CustomErr{
		Message: "hello, alice",
	}
	err = fmt.Errorf("wrap error: %w", err)
	err = fmt.Errorf("second wrap error: %w", err)
	fmt.Println(err)

	fmt.Println(errors.Is(err, CustomErr{}))
	fmt.Println(errors.As(err, &CustomErr{}))

	fmt.Println("---")

	fn := func() (e error) {
		defer func() {
			e = recover().(error)
		}()
		panic(err)
	}
	e := fn()
	fmt.Printf("e: %+v\n", e)
}

func (e CustomErr) Error() string {
	return "custom error"
}
