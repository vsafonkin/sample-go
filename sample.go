package main

import "github.com/vsafonkin/sample-go/web"

func main() {
	if err := web.RunFiberServer("localhost", "8080"); err != nil {
		panic(err)
	}
}
