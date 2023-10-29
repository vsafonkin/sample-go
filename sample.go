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

	var alisa *User
	fmt.Println(alisa == nil)
	fmt.Printf("%p %+[1]v\n", alisa)
}

func doSomething(user *User) {
	user.name = "alisa"
}
