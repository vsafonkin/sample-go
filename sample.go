package main

import (
	"fmt"
	"regexp"
	"time"
)

type User struct {
	name string
	age  int
}

type Tester interface {
	Test()
}

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)
	fmt.Printf("-----\n\n")

	matched, err := regexp.MatchString("l.", "hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)

	reg, err := regexp.Compile("h.?w")
	if err != nil {
		fmt.Println(err)
	}

	result := reg.FindString("hello hew heew")
	fmt.Println(result)
}
