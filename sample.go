package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	ZipCode string `json:"zipcode"`
}

type Users []User

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	users := Users{}
	users.loadUserData("users.json")
	fmt.Printf("%+v\n", users)
}

func (users *Users) loadUserData(path string) error {
	dataFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataFile, users)
	if err != nil {
		return err
	}
	return nil
}
