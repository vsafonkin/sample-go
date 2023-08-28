package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(3)

	done := make(chan struct{})

	go run(50*time.Millisecond, &wg)
	go run(100*time.Millisecond, &wg)
	go run(1000*time.Millisecond, &wg)
	go func() {
		defer close(done)
		fmt.Println("wait all go")
		wg.Wait()
		fmt.Println("all gors are done")
	}()
	<-done

}

func run(timeout time.Duration, wg *sync.WaitGroup) {
	fmt.Println("go start")
	fmt.Println("wait...", timeout, "milliseconds")
	time.Sleep(timeout)
	fmt.Println(timeout, "done")
	wg.Done()
}
