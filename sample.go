package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(123)
	s.Push(456)
	s.Push(789)
	fmt.Println(s.Values())

	n := s.Pop()
	fmt.Println(n, s.Values())

	n = s.Pop()
	fmt.Println(n, s.Values())
}
