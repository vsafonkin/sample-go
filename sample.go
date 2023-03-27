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

type MyString string

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := User{
		name: "Bob",
	}
	bob.sayHello()

	alisa := Admin{}
	err := alisa.sayHello()
	if err != nil {
		fmt.Println(err)
	}

	hello := MyString("hello")
	fmt.Printf("%T\n", hello)
	fmt.Println(hello[0:2])
}

func (u User) sayHello() error {
	if u.name == "" {
		return fmt.Errorf("user name is empty string")
	}
	fmt.Printf("Hello, %s\n", u.name)
	return nil
}
