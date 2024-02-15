package main

import "fmt"

func main() {
	fmt.Println(FindTheNumberPlate(40000))
	fmt.Println(FindTheNumberPlate(1487))
	fmt.Println(FindTheNumberPlate(3))
	fmt.Println(FindTheNumberPlate(17558423))
}

func FindTheNumberPlate(n int) string {
	numPart := (n % 999) + 1
	temp := n / 999
	letters := []rune{'a', 'a', 'a'}
	for i := 0; i < 3; i++ {
		rest := temp % 26
		letters[i] += rune(rest)
		temp = temp / 26
	}
	return fmt.Sprintf("%s%03d", string(letters), numPart)
}
