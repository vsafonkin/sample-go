package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello")
	ctx := context.Background()
	t := reflect.TypeOf(ctx)
	fmt.Println(t.Kind())

	m := map[string]int{
		"hello": 123,
	}
	fmt.Println(reflect.TypeOf(m).Kind())

	ptr := new(string)
	fmt.Println(ptr)

	*ptr = "hello"
	fmt.Println(*ptr)

	fmt.Println(reflect.TypeOf(ptr).Kind())
}
