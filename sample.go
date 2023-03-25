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

	matched, err := regexp.MatchString("a+b=c", "a+b=c")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)

	reg, err := regexp.Compile("h.?w")
	if err != nil {
		fmt.Println(err)
	}

	result := reg.FindAllString("hello hew heew ohowa", -1)
	fmt.Println(result)
}
