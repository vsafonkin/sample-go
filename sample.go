package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

type Day int

const (
	Sunday Day = iota
	Monday
)

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

	fmt.Println(Sunday)
	fmt.Println(Monday)

	printDay(Monday)
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) {
	u.name = name
}

func printDay(day Day) {
	fmt.Println(day)
}
