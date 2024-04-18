package main

import "fmt"

func main() {
	var s []int
	var temp int
	for i := range 8192 {
		s = append(s, i)
		if temp != cap(s) {
			fmt.Println(cap(s))
			temp = cap(s)
		}
	}
}
