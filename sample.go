package main

import (
	"fmt"
)

func main() {
	fmt.Println(wave("hello"))
	fmt.Println(wave("goodbye"))
	fmt.Println(wave(" x yz"))
}

func wave(words string) []string {
	ln := len(words)
	var result []string

	for i := 0; i < ln; i++ {
		if words[i] == ' ' {
			continue
		}
		result = append(result, words[:i]+string(words[i]-32)+words[i+1:])
	}
	return result
}
