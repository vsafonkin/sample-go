package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/common"
)

type Task struct {
	id   int
	done bool
}

type ConsoleLog struct {
	message string
}

func main() {
	task := &Task{
		id: 123,
	}
	log := ConsoleLog{"task is done"}

	doSomething(123, task, log)
}

func doSomething(taskId int, runner common.Runner, printer common.Printer) {
	if err := runner.Run(taskId); err != nil {
		fmt.Println(err)
	}
	printer.Println()
}

func (t *Task) Run(taskId int) error {
	if t.id == taskId {
		t.done = true
		return nil
	}
	return fmt.Errorf("task not matched")
}

func (cl ConsoleLog) Println() {
	fmt.Println(cl.message)
}
