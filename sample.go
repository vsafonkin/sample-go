package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")

	n, ptr := doSomething()
	fmt.Printf("%T, %[1]p, %d\n", ptr, *ptr)
	fmt.Println("n:", n)

	n++
	fmt.Println(*ptr, n)
}

func doSomething() (int, *int) {
	n := 123
	return n, &n
}
