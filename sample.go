package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/mtest"
)

func main() {
	result := mtest.IntSum(123, 231)
	fmt.Println(result)
}
