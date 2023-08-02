package main

import (
	"fmt"
)

type Runner interface {
	Run()
}

type Test struct{}

func main() {
	fmt.Println("hello")

	var test Test
	var runner Runner = test
	fmt.Println(runner == nil)
}

func (t Test) Run() {
	fmt.Println("Test run")
}
