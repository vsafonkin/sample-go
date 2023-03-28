package user

import (
	"fmt"
)

type User struct {
	Name string
	age  int
}

func init() {
	fmt.Println("Init func of user package")
}

func (u User) SayHello() error {
	if u.Name == "" {
		return fmt.Errorf("error! user name is empty")
	}
	fmt.Printf("Hello, %s\n", u.Name)
	return nil
}

func (u *User) SetUserAge(age int) error {
	u.age = age
	return nil
}

func (u User) GetUserAge() int {
	return u.age
}

func (u User) RunUserTest() {
	fmt.Println("RunUserTest method")
}

func (u User) DoNothing() {
	fmt.Println("Do nothing")
}
