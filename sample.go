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
		name: "Bob",
		age:  123,
	}

	fmt.Println(bob.toString())
	bob.setName("Alisa")
	fmt.Println(bob.toString())
	setAge(bob, 456)
	fmt.Println(bob.toString())
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) {
	u.name = name
}

func setAge(user User, age int) {
	user.age = age
}
