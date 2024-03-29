package gendata

import (
	"math/rand"
)

func IntRandomSlice(n int, m int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = rand.Intn(m)
	}
	return out
}

func ByteRandomSlice(n int) []byte {
	out := make([]byte, n)
	for i := range out {
		out[i] = byte(rand.Intn(12345))
	}
	return out
}
