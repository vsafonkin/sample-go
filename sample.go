package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("-----")

	n := 1_000_000
	fmt.Printf("Total time string concat: %v\n", funcExecTimeDuration(runStringConcat, n))
	fmt.Printf("Total time string builder: %v\n", funcExecTimeDuration(runStringBuilder, n))
}

func runStringConcat(n int) {
	var result string
	for i := 0; i < n; i++ {
		result += strconv.Itoa(i)
	}
}

func runStringBuilder(n int) {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString(strconv.Itoa(i))
	}
}

func funcExecTimeDuration(fn func(int), n int) time.Duration {
	start := time.Now()
	fn(n)
	return time.Since(start)
}
