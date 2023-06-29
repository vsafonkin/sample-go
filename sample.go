package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("hello")

	var arr = [100]struct{}{}
	var p unsafe.Pointer
	var st = struct {
		name string
		age  byte
	}{
		name: "Bob",
		age:  123,
	}
	for i := range arr {
		fmt.Println(i)
	}
	fmt.Println("size:", unsafe.Sizeof(arr))
	fmt.Println(st, unsafe.Sizeof(st))
	fmt.Println(p, unsafe.Sizeof(p))
}
