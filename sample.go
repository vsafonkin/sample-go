package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str := "hello"
	sout := unsafe.String(unsafe.StringData(str), 48)
	fmt.Println(sout)
}
