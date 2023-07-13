package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go addIntToArray(arr, i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%p\n", &arr)
	arr = append(arr, 123)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%p\n", &arr)

}

func addIntToArray(arr []int, n int) {
	arr[n] = n
	arr = append(arr, 789)
}
