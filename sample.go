package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("-----")

	if len(os.Args) == 1 {
		fmt.Println("usage: command <domen>")
		os.Exit(0)
	}

	go func() {
		for {
			fmt.Println("gorutinues number:", runtime.NumGoroutine())
			time.Sleep(5 * time.Second)
		}
	}()

	host := os.Args[1]
	var wg sync.WaitGroup
	for i := 1; i <= 65365; i++ {
		wg.Add(1)
		go func(port int) {
			checkTCPConnection(host, port)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done.")
}

func checkTCPConnection(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return
	}
	fmt.Println("connection success:", addr)
	go func() {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error read from connection:", err)
			return
		}
		fmt.Println("read from connection", port, string(buffer[:n]))
	}()
	time.Sleep(5 * time.Second)
	conn.Close()
}
