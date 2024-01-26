package main

import "fmt"

type User struct {
	id   int
	name string
}

type Admin struct {
	User
	access bool
}

func main() {

	alice := Admin{
		User{
			1,
			"Alice",
		},
		true,
	}

	fmt.Printf("%T %+[1]v\n", alice)
	alice.SayHello()
}

func (u User) SayHello() {
	fmt.Println("Hello,", u.name)
}
