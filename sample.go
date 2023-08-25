package main

import (
	"fmt"
	"time"
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

	func() {
		for {
			select {
			case res := <-done:
				fmt.Println("Done chan:", res)
				return
			case v := <-ch:
				fmt.Println("Value chan:", v)
			}
		}
	}()
	fmt.Println("main finish")
}

func run(ch chan<- int, done chan<- bool) {
	defer close(ch)
	fmt.Println("run func")
	ch <- 123
	time.Sleep(300 * time.Millisecond)
	ch <- 456
	time.Sleep(500 * time.Millisecond)
	done <- true
}
