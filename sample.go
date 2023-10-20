package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("-----")
	fmt.Println(123 + 3.45)

	f := 12.3
	res := f + 456.7
	fmt.Printf("%v -> %[1]T\n", res)
}
