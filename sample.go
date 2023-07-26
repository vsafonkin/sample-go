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

	bob := User{
		name: "Bob",
		age:  123,
	}

	sayHello(&bob)
	fmt.Printf("---\n%+v\n", bob)
}

func sayHello(user *User) {
	fmt.Printf("Hello, %s\n", user.name)
	user.name = "Alisa"
	user.age = 456
}
