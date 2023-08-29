package main

import (
	"context"
	"fmt"
	"math/rand"
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
	case v := <-ch:
		fmt.Println("read from ch:", v)
	}
}

func doSomething(ch chan<- int) {
	defer close(ch)
	randomInt := rand.Intn(200)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	ch <- randomInt
}
