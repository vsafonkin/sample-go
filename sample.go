package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	sayHello("Vasya")
}

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}
