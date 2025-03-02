package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"os"
)

type Users []User

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address Address
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	fmt.Println("hello")
	path := "./test_data/users.json"

	users := Users{}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("id: %d, name: %s,\ncity: %s, zipcode: %s\n",
			user.Id,
			user.Name,
			user.Address.City,
			user.Address.Zipcode,
		)
	}
}
