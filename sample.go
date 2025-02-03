package main

import (
	"fmt"
	"github.com/vsafonkin/sample-go/web"
	"log"
)

func main() {
	fmt.Println("Hello, go")
	if err := web.RunGinServer("localhost", "81"); err != nil {
		log.Fatal()
	}
}
