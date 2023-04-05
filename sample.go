package main

import "fmt"

type Day int

const (
	_ Day = iota
	Sunday
	Monday
)

type Username string

func main() {
	fmt.Println("-----")
	showDayValue(Sunday)
	sayHello("Bob")
	var alisa Username = "Alisa"
	fmt.Println(alisa + "bob")
}

func showDayValue(d Day) {
	fmt.Printf("Day: %d\n", d)
}

func sayHello(name Username) {
	fmt.Printf("Hey, %s\n", name)
}
