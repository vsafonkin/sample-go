package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	bob := User{
		name: "Bob",
		age:  123,
	}

	fmt.Println(bob.toString())
	bob.setName("Alisa")
	fmt.Println(bob.toString())
	setAge(&bob, 456)
	fmt.Println(bob.toString())

	nums := []int{1, 2, 3}
	setFirstSliceItem(nums, 789)
	fmt.Println(nums)

	test := make([]int, 10)
	test = append(test, 1)
	fmt.Println(len(test), test)

	clubs := map[string]string{
		"real": "madrid",
		"psg":  "paris",
	}
	fmt.Println(clubs)
	setFirstMapValue(clubs, "barselona")
	fmt.Println(clubs)
}

func (u User) toString() string {
	return fmt.Sprintf("User name: %s, User age: %d\n", u.name, u.age)
}

func (u *User) setName(name string) {
	u.name = name
}

func setAge(user *User, age int) {
	user.age = age
}

func setFirstSliceItem(slice []int, newValue int) {
	slice[0] = newValue
}

func setFirstMapValue(mMap map[string]string, newValue string) {
	for k, _ := range mMap {
		mMap[k] = newValue
		break
	}
}
