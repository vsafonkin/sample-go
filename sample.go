package main

import (
	// "bufio"
	"fmt"
	// "net/http"
	// _ "net/http/pprof"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:8080", nil))
	// }()

	fmt.Println("hello")
	doSomething()

	// buf := bufio.NewReader(os.Stdin)
	// fmt.Println("Press enter to finish program")
	// buf.ReadBytes('\n')
}

func doSomething() {
	fmt.Println("do something")
}
