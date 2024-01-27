package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/gendata"
)

func main() {

	bs := gendata.ByteRandomSlice(1000)
	fmt.Println(string(bs))
}
