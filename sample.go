package main

import "github.com/vsafonkin/sample-go/web"

func main() {
	web.RunFiberServer("localhost", "8080")
}
