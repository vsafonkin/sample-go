package prime

import (
	"testing"
)

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

func TestIsPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"is prime", args{127}, true},
		{"not prime", args{99}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
