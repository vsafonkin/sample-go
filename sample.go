package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	names := []string{"bob", "alisa", "mike", "test"}
	for i, v := range names {
		switch v {
		case "bob":
			fmt.Println(i, "Bob case")
		case "alisa":
			fmt.Println(i, "Alisa case")
		case "mike":
			fmt.Println(i, "Mike case")
		default:
			fmt.Println(i, "Default case")
		}
	}
}
