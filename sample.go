package main

import (
	"fmt"
	"slices"
)

type User struct {
	id   int
	name string
}

type Users []User

func main() {

	slc := []int{3, 5, 2, 7, 9, 3, 1, 5, 4, 4, 6}
	slices.Sort(slc)
	fmt.Println(slc)

	str := []string{"hello", "goodbye", "alice", "bob"}
	slices.Sort(str)
	fmt.Println(str)

	users := Users{
		{id: 123, name: "bob"},
		{id: 456, name: "alice"},
		{id: 789, name: "mike"},
		{id: 111, name: "aloha"},
		{id: 123, name: "bob"},
	}

	fmt.Println(users)
	slices.SortFunc(users, func(a User, b User) int {
		out := 0
		if a.name < b.name {
			out = -1
		}
		if a.name > b.name {
			out = 1
		}
		return out
	})
	fmt.Println(users)
}
