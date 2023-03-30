package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	var wg sync.WaitGroup
	num := 10
	wg.Add(num)

	out := make(chan int, num)
	for i := 0; i < num; i++ {
		id := i
		go func() {
			defer wg.Done()
			rnum := rand.Intn(1000)
			time.Sleep(time.Duration(rnum) * time.Millisecond)
			fmt.Printf("Close grt %d, interval %d ms\n", id, rnum)
			out <- rnum
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	fmt.Printf("Result: %v\n", result)
}
