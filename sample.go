package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/vsafonkin/sample-go/tcpserver"
)

func main() {
	fmt.Println("---")
	if len(os.Args) != 3 {
		fmt.Println("usage: sample <address> <port>")
		os.Exit(0)
	}

	ctx := context.Background()
	done := make(chan struct{})
	server := tcpserver.NewTCPServer(ctx, os.Args[1], os.Args[2])
	go func() {
		err := server.Run()
		if err != nil {
			log.Println("run server error:", err)
			return
		}
		<-done
	}()
	time.Sleep(60 * time.Second)
	err := server.Stop()
	if err != nil {
		log.Println("stop server error:", err)
	}
	done <- struct{}{}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
