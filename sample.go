package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func main() {
	fmt.Println("hello")
	doSomething()

	bob := newUser("Bob", 123)
	fmt.Println(bob)
}

func doSomething() {
	fmt.Println("func doSomething")
}

func newUser(name string, age int) User {
	return User{
		name: name,
		age:  age,
	}
}
