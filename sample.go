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

	b := <-ch
	fmt.Println(b)
	fmt.Println("end")
}
