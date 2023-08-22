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

	ch := make(chan string)
	go run(ch)
	time.Sleep(200 * time.Millisecond)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	for v := range ch {
		fmt.Printf("Value: %s\n", v)
	}
}

func run(ch chan<- string) {
	defer close(ch)
	fmt.Println("run func")
	ch <- "hello"
	ch <- "goodbye"
	ch <- "12345"
}
