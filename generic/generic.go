package generic

import (
	"fmt"
)

func HandleMap[K comparable, V any](m map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func HandleSlice[T any, V any](arg T, slc []V) {
	fmt.Println("First arg:", arg)
	fmt.Println("Slice:", slc)
}

func Run() error {
	m := map[string]int{}
	m["hello"] = 123
	m["goodbye"] = 456
	keys, values := HandleMap(m)
	fmt.Println("Keys slice:", keys)
	fmt.Println("Values slice:", values)
	fmt.Println("-----")

	slc := []bool{true, false, true}
	slc_2 := []int{7, 11, 13}
	HandleSlice("Hello", slc)
	HandleSlice(1.2e3, slc_2)
	return nil
}
