package main

import (
	"fmt"
	"testing"
)

func AddItems() {
	var data = []string{}
	for i := 0; i <= 65600; i++ {
		data = append(data, fmt.Sprintf("%d", i))
	}
}
func AddItems2() {
	data := make([]string, 65600)
	for i := 0; i <= 65600; i++ {
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

func BenchmarkSomething(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		func(n int) int {
			return i * i * i
		}(i)
	}
	b.ReportAllocs()
}
