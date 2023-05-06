package main

import (
	"fmt"
)

type Runner struct {
	id int
}

func main() {
	fmt.Println("-----")

	var runner *Runner
	fmt.Println(runner, runner == nil)
	runner.run()
}

func (r *Runner) run() error {
	fmt.Println("Run...")
	return nil
}
