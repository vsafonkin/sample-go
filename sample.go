package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")

	ptr := doSomething()
	fmt.Printf("%T, %[1]p, %d\n", ptr, *ptr)
}

func doSomething() *int {
	n := 123
	return &n
}
