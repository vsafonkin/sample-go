package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func main() {
	fmt.Println("hello")

	var bob *User
	fmt.Println(bob == nil)
	bob.test()
}

func (u *User) test() {
	fmt.Println("User method: test")
}
