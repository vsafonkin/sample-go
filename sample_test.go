package main

import (
	"fmt"
	"testing"
)

func AddItems() {
	var data = []string{}
	for i := 0; i <= 65550; i++ {
		data = append(data, fmt.Sprintf("%d", i))
	}
}
func AddItems2() {
	data := make([]string, 65550)
	for i := 0; i <= 65550; i++ {
		data = append(data, fmt.Sprintf("%d", i))
	}
}

func BenchmarkAddItems(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		AddItems()
	}
	b.ReportAllocs()
}
func BenchmarkAddItems2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		AddItems2()
	}
	b.ReportAllocs()
}

func ByteSlice() *[]byte {
	bs := new([]byte)
	temp := make([]byte, 2048)
	*bs = temp
	return bs
}

func BenchmarkByteSlice(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		ByteSlice()
	}
	b.ReportAllocs()
}
