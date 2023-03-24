package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

type Admin struct {
	name string
}

type Tester interface {
	Test()
}

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := User{
		name: "Bob",
		age:  123,
	}

	var tester Tester = bob
	runTest(tester)

	sayHello(tester.(User))

}

func runTest(tester Tester) error {
	tester.Test()
	return nil
}

func (u User) Test() {
	fmt.Println("User test method")
}

func (u Admin) Test() {
	fmt.Println("Admin test method")
}

func sayHello(user User) {
	fmt.Printf("Hello, %s\n", user.name)
}
