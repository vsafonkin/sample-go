package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/prime"
)

func main() {
	fmt.Println("---")
	n := 100_000_000
	slc := prime.PrimeSlice(n)
	fmt.Println(len(slc))
}
