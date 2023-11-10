package prime

import "testing"

func BenchmarkPrimeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeSlice(1_000_000)
		b.ReportAllocs()
	}
}

func BenchmarkConcPrimeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcPrimeSlice(1_000_000)
		b.ReportAllocs()
	}
}
