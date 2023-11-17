package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str := "hello"
	sh := unsafe.String(unsafe.StringData(str), 48000)
	fmt.Println(sh)
}
