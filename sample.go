package main

import (
	"fmt"
	"github.com/pkg/profile"
)

var test = func() int {
	fmt.Println("test func")
	return 123
}()

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")
}
