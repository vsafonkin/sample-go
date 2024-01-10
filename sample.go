package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----")

	var count int
	for {
		fmt.Println(count)
		count++
		time.Sleep(1 * time.Second)
	}
}
