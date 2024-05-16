package main

import (
	"github.com/vsafonkin/sample-go/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(123)
	s.Push(456)
	s.Push(789)
}
