package main

import (
	"fmt"
	"github.com/vsafonkin/sample-go/progressbar"
)

func main() {
	fmt.Println("hello")

	done := make(chan struct{})
	go progressbar.ProgressBar(done)
	<-done
	fmt.Println("finish")
}
