package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		main()
	}
}

func BenchmarkFib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fib(50)
	}
}
