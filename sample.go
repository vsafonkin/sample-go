package main

import (
	"fmt"
)

type Fn func()

func main() {
	fmt.Println("-----")

	DoSomething, counter := func() (Fn, *int) {
		var counter int
		return func() {
			counter++
			DoSomething()
		}, &counter
	}()
	fmt.Println(*counter)
	DoSomething()
	fmt.Println(*counter)
	DoSomething()
	DoSomething()
	fmt.Println(*counter)
}

func DoSomething() {
	fmt.Println("Do something")
}
