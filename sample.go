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
		ch <- time.Now()
	}()

	b := <-ch
	fmt.Println(b)
	fmt.Println("end")
}
