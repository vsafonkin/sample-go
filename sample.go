package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("-----")

	doSomething("hello", "goodbye", "gogogo")

	names := []string{"bob", "alisa", "mike"}
	doSomething(names...)
}

func doSomething(arr ...string) {
	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}
}
