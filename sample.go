package main

import (
	"fmt"
)

type Counter interface {
	count() int
}

func main() {
	fmt.Println("hello")
}
