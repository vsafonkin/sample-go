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
	err := bob.setName("Alisa")
	if err != nil {
		fmt.Println("Set name error:", err)
	}
	fmt.Println(bob.toString())
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) error {
	if u == nil {
		return fmt.Errorf("object user is nil")
	}
	u.name = name
	return nil
}
