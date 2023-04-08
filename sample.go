package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")

	f := 12.345
	fmt.Printf("%f %[1]e\n", f)
}

func DoSomething() {
	fmt.Println("Do something")
}
