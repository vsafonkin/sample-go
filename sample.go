package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("-----")

	f := 12.345
	fmt.Printf("%f %[1]e\n", f)

	var p *int
	fmt.Println(p)

	t := reflect.TypeOf(p)

	fmt.Println(t.Name())
}

func DoSomething() {
	fmt.Println("Do something")
}
