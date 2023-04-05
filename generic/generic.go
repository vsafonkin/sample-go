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

func Run() error {
	m := map[string]int{}
	m["hello"] = 123
	m["goodbye"] = 456
	keys, values := HandleMap(m)
	fmt.Println("Keys slice:", keys)
	fmt.Println("Values slice:", values)
	return nil
}
