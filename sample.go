package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
	nums []int
}

func main() {
	fmt.Println("hello")

	bob := User{
		name: "Bob",
		age:  123,
		nums: []int{1, 2, 3},
	}

	var alisa User

	fmt.Println(bob)
	fmt.Println(alisa, alisa.nums == nil)

	mike := User{"Mike", 456, []int{4, 5, 6}}
	fmt.Println(mike)

	anonim := struct {
		name string
	}{
		name: "John",
	}

	fmt.Println(anonim)
}
