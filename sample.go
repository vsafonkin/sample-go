package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func main() {

	bob := User{
		name: "bob",
	}
	doSomething(&bob)

	fmt.Printf("%+v\n", bob)
}

func doSomething(user *User) {
	user.name = "alisa"
}
