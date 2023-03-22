package main

import (
	"fmt"
	"reflect"
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
	fmt.Println("-----")
	out := anyType(bob)
	fmt.Println(reflect.TypeOf(out), out)
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

func anyType(obj interface{}) interface{} {
	fmt.Println(reflect.TypeOf(obj), obj)
	type Sample struct {
		a int
	}
	var sample Sample
	sample.a = 123
	return sample
}
