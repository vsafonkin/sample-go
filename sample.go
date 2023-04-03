package main

import (
	"fmt"
	"time"

	"github.com/vsafonkin/sample-go/stdlib_json"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	usersData := stdlib_json.Users{}
	err := usersData.LoadUserData("./test_data/users.json")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("Users data: %+v\n", usersData)
}
