package main

import (
	"fmt"
)

type Fn func(str string) int

func main() {
	fmt.Println("-----")

	doSomething, counter := initFnCounter(func(str string) int { return len(str) })
	fmt.Println(*counter)
	fmt.Printf("Len: %d\n", doSomething("hello"))
	fmt.Printf("Len: %d\n", doSomething("goodbye"))
	fmt.Printf("Len: %d\n", doSomething("aloha"))
	fmt.Println(*counter)
	fmt.Printf("Len: %d\n", doSomething("golang"))
	fmt.Println(*counter)

}

func initFnCounter(fn Fn) (Fn, *int) {
	var counter int
	return func(str string) int {
		counter++
		return fn(str)
	}, &counter
}
