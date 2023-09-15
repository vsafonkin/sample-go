package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	dat, err := os.Open("./test_data/users.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(dat)
}
