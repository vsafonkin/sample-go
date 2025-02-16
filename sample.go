package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"os"
)

type Users []User

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	path := "./test_data/users.json"
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("defer")
	}()

	users := Users{}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)

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

	var users2 Users
	err = json.NewDecoder(f).Decode(&users2)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users2 {
		fmt.Printf("id: %d, name: %s\n", user.Id, user.Name)
	}
}
