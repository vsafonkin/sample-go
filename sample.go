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

	for v := range ch {
		fmt.Println(v)
	}
}

func run(id int, ch chan<- int) {
	i := id
	go func() {
		ch <- id
		if i == 789 {
			time.Sleep(2 * time.Second)
			close(ch)
		}
		fmt.Printf("End goroutine id: %d\n", id)
	}()
}
