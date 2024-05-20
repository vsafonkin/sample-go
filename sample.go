package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	n := 100_000_000
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(70)
	}
	res1 := make([]int, n)
	go func() {
		defer wg.Done()
		start := time.Now()
		for i := range s {
			res1[i] = Fib(s[i])
		}
		d := time.Since(start)
		fmt.Println("time 1:", d)
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		_ = Mem(s)
		d := time.Since(start)
		fmt.Println("time 2:", d)
	}()
	wg.Wait()
}

func Fib(n int) int {
	a, b, c := 0, 1, 0
	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func Mem(s []int) []int {
	out := make([]int, len(s))
	m := make(map[int]int, len(s))
	for i := range s {
		if v, ok := m[s[i]]; ok {
			out[i] = v
			continue
		}
		r := Fib(s[i])
		out[i] = r
		m[s[i]] = r
	}
	return out
}
