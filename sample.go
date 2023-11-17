package main

import (
	"fmt"
	"unsafe"
)

func main() {

	s := []int{1, 2, 3}
	sp := (*[3]int)(unsafe.Pointer(&s))
	fmt.Println(*sp)
}
