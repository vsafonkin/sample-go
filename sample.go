package main

import (
	"fmt"
)

type Runner interface {
	Run() error
}

type Test int

func main() {
	fmt.Println("-----")

	var runner Test
	err := Exec(runner)
	fmt.Println(err)
}

func Exec(runner Runner) error {
	return runner.Run()
}

func (t Test) Run() error {
	fmt.Println("Test m")
	return nil
}
