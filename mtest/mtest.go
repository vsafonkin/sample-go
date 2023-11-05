package mtest

import (
	"fmt"
	"sync"
	"time"
)

type Runner interface {
	Run(n int) error
}

func HelloRunner(n int, runner Runner) {
	err := runner.Run(n)
	if err != nil {
		fmt.Println(err)
	}
}

func IntSum(a, b int) int {
	return a + b
}

func doSomething() string {
	return "do something"
}

func dataRace() int {
	var counter int
	go func() {
		counter++
	}()
	return counter
}

func RunSomething(td time.Duration) int {
	var wg sync.WaitGroup
	var counter int
	ch := make(chan int)
	done := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case ch <- counter:
				counter++
			}
		}
	}()
	go func() {
		timeout := time.After(td)
		for {
			select {
			case <-ch:
			case <-timeout:
				done <- struct{}{}
			}
		}
	}()
	wg.Wait()
	return counter
}
