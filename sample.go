package main

import (
	"fmt"
	"time"

	nilvalue "github.com/vsafonkin/sample-go/nil"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	nilvalue.ZeroValues()
	fmt.Println("-----")
	nilvalue.NilValues()
	fmt.Println("-----")
	nilvalue.Run()
}
