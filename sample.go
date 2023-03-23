package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

type Tester interface {
	Test()
}

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")
}
