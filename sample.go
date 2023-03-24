package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
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
	fmt.Println(bob)

	var tester Tester = bob
	runTest(tester)
	i := tester.(User)
	fmt.Println(i)

}

func runTest(tester Tester) error {
	tester.Test()
	return nil
}

func (u User) Test() {
	fmt.Println("User test method")
}
