package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer test("main")
	fmt.Println("hello")
	go func() {
		defer test("anonim grt")
		log.Printf("%s\n", "goroutine")
	}()
	time.Sleep(100 * time.Millisecond)
}

func test(str string) {
	log.Printf("%s: %s\n", "defer", str)
}
