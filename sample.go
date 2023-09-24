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
	fmt.Println("goodbye")
}
