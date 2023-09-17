package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	n := 123
	done := make(chan struct{})

	go func() {
		n++
		done <- struct{}{}
	}()

	<-done
	fmt.Println(n) // race conditions
	time.Sleep(10 * time.Millisecond)
	fmt.Println(n)
}
