package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	name string
	age  int
}

type sum func(int, int) int

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime)

	bnum := 0b11111110
	fmt.Println(bnum)

	xnum := 0xFE
	fmt.Println(xnum)

	unicode := '\u0061'
	fmt.Println(unicode)

	var f float64 = 254.1
	result := int(f)
	fmt.Println(result)
	fmt.Printf("Type: %T\nBinary: %b\nHex: %X\n", result, result, result)
	fmt.Printf("%d\n", 0b11111011)

	myFunc := test
	myFunc("hello", "goodbye")

	err := runSomeFunc(myFunc, "abrakadabra")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	sep()

	userNames := getUserNames(User{name: "bob", age: 123}, User{name: "alisa", age: 456})
	fmt.Println(userNames)

	testArgs("hello", "goodbye")
	emptyFunc()

	var sFunc sum
	fmt.Println(reflect.TypeOf(sFunc))

}

func add(a int, b int) int {
	return a + b
}

func test(strarray ...string) error {
	for _, value := range strarray {
		fmt.Println(value)
	}
	return nil
}

func runSomeFunc(someFunc func(...string) error, str string) error {
	err := someFunc(str)
	return err
}

func sep() {
	fmt.Printf("\n-----\n\n")
}

func getUserNames(users ...User) []string {
	var userNames []string
	for _, v := range users {
		userNames = append(userNames, v.name)
	}
	return userNames
}

func testArgs(...string) {
	fmt.Println("testArgs")
}

func emptyFunc() (greeting string) {
	greeting = "hello"
	return
}
