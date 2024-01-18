package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(fmt.Sprintf("./%s/%s", "pprof", *cpuprofile))
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ch := make(chan int)
	go func() {
		time.Sleep(101 * time.Millisecond)
		ch <- doSomething()
	}()

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(100 * time.Second)
		timeout <- true
	}()

	select {
	case n := <-ch:
		fmt.Println("recieved:", n)
	case <-timeout:
		fmt.Println("timeout")
	}
}

func doSomething() int {
	return 42
}
