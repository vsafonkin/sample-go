package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-----")

	fmt.Println(os.Getenv("HOME"))
}
