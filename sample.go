package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vsafonkin/sample-go/tcpserver"
)

func main() {
	fmt.Println("---")
	if len(os.Args) != 3 {
		fmt.Println("usage: sample <address> <port>")
		os.Exit(0)
	}

	server := tcpserver.NewTCPServer(os.Args[1], os.Args[2])
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
