package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----")
	s1 := doSomething()
	println(s1)
}

//go:noinline
func doSomething() [][]int {

	t := make([]int, 1000)
	n := 1000
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
		t[i] = i
	}
	return [][]int{t, s, run()}
}

//go:noinline
func run() []int {
	n := 5000
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}
