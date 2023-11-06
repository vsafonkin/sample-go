package main

import (
	"fmt"
	"time"

	"github.com/vsafonkin/sample-go/mtest"
)

func main() {
	result := mtest.RunSomething(1 * time.Second)
	fmt.Println(result)
}
