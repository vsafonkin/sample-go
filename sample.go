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

	intMax := int(^uint(0) >> 1)
	fmt.Printf("%v %08[1]b -> %[1]T\n", intMax)
	fmt.Println(intMax + 12345)

	var t byte = 253
	sum := t + 123
	fmt.Printf("-----\n%v -> %[1]T\n", sum)
}
