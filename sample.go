package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 123
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("goodbye")
}
