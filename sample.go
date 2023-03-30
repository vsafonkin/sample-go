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
	wg.Add(10)

	for i := 0; i < 10; i++ {
		id := i
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Close grt %d\n", id)
		}()
	}
	wg.Wait()
}
