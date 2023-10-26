package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a = 231
	var b = 253
	var c = 123
	var sum = a + b + c
	fmt.Printf("Size: %d, Pointer: %p\n", unsafe.Sizeof(sum), &sum)
	doSomething()

	str := "goodbye"
	fmt.Println(str)
}

func doSomething() {
	message := "hibob"
	userId := 123321
	fmt.Println("start...", message, &message, userId, &userId)
}
