package main

import (
	"fmt"
)

type Runner interface {
	Run()
}

func main() {
	fmt.Println("-----")

	var runner Runner
	fmt.Println(runner)
}
