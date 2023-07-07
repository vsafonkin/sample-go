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

	bob := User{
		name: "Bob",
		age:  123,
	}

	var alisa User

	fmt.Println(bob)
	fmt.Println(alisa)

	mike := User{"Bob", 123}
	fmt.Println(mike == bob)

	anonim := struct {
		name string
	}{
		name: "John",
	}

	fmt.Println(anonim)
}
