package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/user"
)

func main() {

	bob := user.User{
		Name: "Bob",
	}
	bob.SetUserAge(123)

	fmt.Printf("%#v\n", bob)
}
