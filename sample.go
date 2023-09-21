package main

import (
	"fmt"
	"github.com/pkg/profile"
	"math/rand"
)

func main() {
	prof := profile.Start()
	defer prof.Stop()
	fmt.Println("hello")
	doSomething()
}

func doSomething() {
	fmt.Println("do something")
	for {
		if rand.Intn(5_000_000_000) == 123456789 {
			fmt.Println("Done")
			break
		}
	}
}
