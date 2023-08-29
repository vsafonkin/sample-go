package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan int)
	go run(ctx, ch)
	go doSomething(ch)
	time.Sleep(100 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func run(ctx context.Context, ch <-chan int) {
	select {
	case <-ctx.Done():
		fmt.Println("run func is done")
	case <-ch:
		fmt.Println("read from ch")
	}
}

func doSomething(ch chan<- int) {
	defer close(ch)
	time.Sleep(150 * time.Millisecond)
	ch <- 123
}
