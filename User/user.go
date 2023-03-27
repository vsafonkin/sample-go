package user

import (
	"fmt"
)

type User struct {
	Name string
}

func (u User) SayHello() error {
	if u.Name == "" {
		return fmt.Errorf("error! user name is empty")
	}
	fmt.Printf("Hello, %s\n", u.Name)
	return nil
}
