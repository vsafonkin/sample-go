package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Persistence(39))
}

func Persistence(n int) int {
	var counter, acc = 1, 1
	s := fmt.Sprintf("%d", n)
	if len(s) == 1 {
		return 0
	}
	for _, v := range s {
		d, _ := strconv.Atoi(string(v))
		acc *= d
	}
	counter += Persistence(acc)
	return counter
}
