package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	n := 123

	go func() {
		n++
	}()

	fmt.Println(n) // race conditions
	time.Sleep(10 * time.Millisecond)
	fmt.Println(n)
}
