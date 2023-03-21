package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

type Person struct {
	User
	admin bool
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

	alisa := Person{
		User: User{
			name: "Alisa",
			age:  22,
		},
		admin: true,
	}

	fmt.Println(alisa.toString())
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) {
	u.name = name
}
