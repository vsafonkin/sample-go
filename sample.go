package main

import (
	"fmt"
	"time"

	"github.com/vsafonkin/sample-go/user"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := user.User{
		Name: "Bob",
	}

	bob.SayHello()
	fmt.Println(bob)

	bob.SetUserAge(123)
	fmt.Println(bob)

	var alisa user.User
	fmt.Println(alisa)
	alisa.Name = "Alisa"
	alisa.SetUserAge(456)
	fmt.Println(alisa)
}
