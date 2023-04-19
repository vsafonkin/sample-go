package main

import (
	"fmt"
)

type Runner interface {
	Run() error
}

type Tester interface {
	StartTest() error
}

type Process struct {
	id int
	Runner
	Tester
}

func main() {
	fmt.Println("-----")

	var p Process
	fmt.Printf("id: %d\n", p.id)
	err := Exec(p)
	fmt.Println(err)
}

func Exec(runner Runner) error {
	return runner.Run()
}

func (a Process) Run() error {
	fmt.Println("Test m")
	return nil
}
