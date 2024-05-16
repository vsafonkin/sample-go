package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		main()
	}
}
