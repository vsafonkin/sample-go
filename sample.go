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
}
