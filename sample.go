package main

import (
	"fmt"
)

type TestError struct {
	Message   string
	ErrStatus int
}

func (e TestError) Error() string {
	return fmt.Sprintf("error message: %s, error status: %d\n", e.Message, e.ErrStatus)
}

func main() {
	fmt.Println("-------------")

	ch := make(chan int)
	done := make(chan bool)
	go run(ch, done)

	select {
	case res := <-done:
		fmt.Println("Done chan:", res)
	case v := <-ch:
		fmt.Println("Value chan:", v)
	}
	fmt.Println("main finish")
}

func run(ch chan<- int, done chan<- bool) {
	defer close(ch)
	fmt.Println("run func")
	ch <- 123
}
