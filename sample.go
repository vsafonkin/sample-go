package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	names := []string{"bob", "alisa", "mike"}
	test(names...)
	test("hello", "goodbye")
}

func test(names ...string) {
	for i, v := range names {
		fmt.Printf("%d: %s\n", i, v)
	}
}
