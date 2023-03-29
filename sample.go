package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	ch := make(chan time.Time)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- time.Now()
	}()

	ich := make(chan int)
	run(ich)
	b := <-ch
	fmt.Println(b)

	i := <-ich
	fmt.Println(i)
	fmt.Println("end")
}

func run(ch chan<- int) {
	go func() {
		ch <- 123
	}()
}
