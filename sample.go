package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

type Admin struct {
	User
	age int
}

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := User{
		name: "Bob",
	}
	bob.sayHello()
}

func (u User) sayHello() error {
	if u.name == "" {
		return fmt.Errorf("User name is empty string")
	}
	fmt.Printf("Hello, %s\n", u.name)
	return nil
}
