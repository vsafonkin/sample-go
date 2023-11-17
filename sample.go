package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str := "hello"
	p := (*[2]int)(unsafe.Pointer(&str))
	p[1] = 48000
	fmt.Println(str)
}
