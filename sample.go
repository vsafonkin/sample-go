package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	fmt.Println(getRandomInt(100))

}

func getRandomInt(n int) int {
	done := make(chan struct{})
	ch := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			select {
			case ch <- rand.Intn(1000):
			case <-done:
			}
		}()
	}
	result := <-ch
	close(done)
	return result
}
