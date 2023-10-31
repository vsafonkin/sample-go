package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	c int
}

func main() {
	counter := Counter{
		c: 0,
	}

	go func() {
		defer counter.Unlock()
		counter.Lock()
		for i := 0; i < 100000; i++ {
			counter.c = i
		}
	}()

	time.Sleep(10 * time.Nanosecond)
	counter.Lock()
	fmt.Printf("Counter: %d, goroutines: %d\n", counter.c, runtime.NumGoroutine())
	counter.Unlock()
}
