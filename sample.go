package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	prof := profile.Start()
	defer prof.Stop()

	fmt.Println("hello")
}
