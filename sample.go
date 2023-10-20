package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")
	str := "-----\n\u222F\n-----"
	fmt.Println(str)
}
