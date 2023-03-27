package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

}

func (u User) SayHello() error {
	if u.name == "" {
		return fmt.Errorf("error! user name is empty")
	}
	fmt.Printf("Hello, %s\n", u.name)
	return nil
}
