package main

import "fmt"

func main() {
	n := 10
	done := make(chan int)
	var counter int
	for i := range 10 {
		go func() {
			done <- i
		}()
	}

	for v := range done {
		fmt.Println(v)
		counter++
		if counter == n {
			close(done)
		}
	}
}
