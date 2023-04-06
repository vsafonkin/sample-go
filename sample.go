package main

import (
	"fmt"
)

type Fn func(str string) int

func main() {
	fmt.Println("-----")

	var arr []int
	fmt.Println(arr == nil)

	doSomething("hello", func(str string) int { return len(str) })
}

func doSomething(str string, fn Fn) {
	fmt.Println(fn(str))
}
