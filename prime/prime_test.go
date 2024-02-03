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

func Test_doSomething(t *testing.T) {
	type args struct {
		n int
		s string
		b bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		{"sample", args{123, "Alice", true}, 5, "done", false},
		{"error", args{321, "fail", true}, 0, "", true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := doSomething(tt.args.n, tt.args.s, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("doSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("doSomething() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("doSomething() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
