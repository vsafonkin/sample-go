package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
}
