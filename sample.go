package main

import "fmt"

type Day int

const (
	_ Day = iota
	Sunday
	Monday
)

func main() {
	fmt.Println("-----")
	showDayValue(Sunday)
}

func showDayValue(d Day) {
	fmt.Printf("Day: %d\n", d)
}
