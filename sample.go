package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := User{
		name: "bob",
		age:  123,
	}

	fmt.Println(bob)
	sayHello(bob)
	fmt.Println(bob)
}

func sayHello(user User) error {
	fmt.Printf("Hello %s\n", user.name)
	user.name = "alisa"
	return nil
}
