package main

import (
	"fmt"
)

type User struct {
	name string
}

func main() {
	fmt.Println("hello")
	doSomething()

	stc := new(struct {
		name string
	})
	*stc = User{
		name: "bob",
	}
	fmt.Println(stc)
}

func doSomething() {
	fmt.Println("do something")
}
