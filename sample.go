package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var once sync.Once

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("once function")
			})
		}()
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("done")
}
