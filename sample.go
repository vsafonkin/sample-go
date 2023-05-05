package main

import (
	"fmt"
)

type Runner struct{}

func main() {
	fmt.Println("-----")

	var runner *Runner
	fmt.Println(runner == nil)
	runner.run()
}

func (r *Runner) run() error {
	fmt.Println("Run...")
	return nil
}
