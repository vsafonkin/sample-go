package main

import (
	"fmt"
	"github.com/vsafonkin/sample-go/web"
	"log"
)

func main() {
	fmt.Println("Hello, go")
	if err := web.RunGinServer("localhost", "8080"); err != nil {
		log.Fatal()
	}
}
