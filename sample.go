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
		name: "bob",
		age:  123,
	}

	fmt.Println(bob)
	sayHello(bob)
	fmt.Println(bob)

	sampleMap := map[int]string{
		0: "hello",
	}
	fmt.Println(sampleMap)
	modMap(sampleMap)
	fmt.Println(sampleMap)
}

func sayHello(user User) error {
	fmt.Printf("Hello %s\n", user.name)
	user.name = "alisa"
	return nil
}

func modMap(m map[int]string) {
	fmt.Println("hello => goodbye")
	m[0] = "goodbye"
}
