package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	message string
	code    int
}

func main() {
	err := errors.New("test error")
	fmt.Println(err)
	println(err)

	var e error
	println(e)

	var ce *CustomErr
	println(ce)

	e = ce
	println(e)
}

func (ce *CustomErr) Error() string {
	return ce.message
}
