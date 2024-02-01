package common

type Runner interface {
	Run(taskId int) error
}

type Printer interface {
	Println()
}
