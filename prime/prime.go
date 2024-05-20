package prime

import (
	"fmt"
	"sync"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return isPrime
}

func PrimeSlice(n int) []int {
	var output []int
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			output = append(output, i)
		}
	}
	return output
}

func ConcPrimeSlice(n int) []int {
	var output []int
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if IsPrime(i) {
				ch <- i
			}
		}(i)
	}

	go func() {
		for v := range ch {
			output = append(output, v)
		}
	}()

	wg.Wait()
	close(ch)
	return output
}
