package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fib(50))
}

func Fib(n int) []int {
	var out []int
	a, b, c := 0, 0, 1
	for i := 0; i < n; i++ {
		a = b + c
		b = c
		c = a
		out = append(out, c)
	}
	return out
}
