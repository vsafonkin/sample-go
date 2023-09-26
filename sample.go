package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")

	fmt.Println(sum(1, 2, 3))

	nums := []int{4, 5, 6}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
