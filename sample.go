package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	ch := make(chan int)
	run(123, ch)
	run(456, ch)
	run(789, ch)

	fmt.Println(<-ch)
	time.Sleep(2 * time.Second)
}

func run(id int, ch chan<- int) {
	go func() {
		ch <- id
		fmt.Printf("End goroutine id: %d\n", id)
	}()
}
