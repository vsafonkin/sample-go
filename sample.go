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
	fmt.Println(bob.toString())
	bob.setName("Alisa")
	fmt.Println(bob.toString())

	runTest(bob)

	data := map[string]interface{}{
		"a": "hello",
		"b": 123,
	}
	data["c"] = true
	fmt.Println(data)
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) {
	u.name = name
}

func (u User) Test() {
	fmt.Println("User Test")
}

func runTest(test Tester) error {
	test.Test()
	return nil
}
