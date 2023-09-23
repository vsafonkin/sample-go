package main

import (
	"fmt"
	"github.com/pkg/profile"
	"io"
	"os"
)

func main() {
	prof := profile.Start()
	defer prof.Stop()

	io.WriteString(os.Stdout, "hello\n")

	arrFn := make([]func(), 5)
	for i := range [5]struct{}{} {
		arrFn[i] = func() {
			fmt.Println(i)
		}
	}

	arrFn[2]()

}
