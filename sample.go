package main

import (
	"fmt"
	"github.com/pkg/profile"
	"io"
	"os"
)

func main() {
	prof := profile.Start(profile.ProfilePath("./pprof"))
	defer prof.Stop()

	fmt.Println("hello")
	io.WriteString(os.Stdout, "-----\n")
}
