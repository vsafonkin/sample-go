package main

import (
	"fmt"
	"github.com/vsafonkin/sample-go/web"
	"log"
)

func main() {
	fmt.Println("Hello, go")
	if err := web.RunGinServer("0.0.0.0", "8083"); err != nil {
		log.Fatal()
	}
}
