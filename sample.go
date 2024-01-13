package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("-----")

	f, err := os.Open("./sample_test.go")
	if err != nil {
		panic(fmt.Errorf("open file error: %v", err))
	}

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			break
		}
		fmt.Println(string(filterRune(buf[:n], []rune{' ', '\t', '\n'})))

	}
}

func filterRune(slc []byte, r []rune) []byte {
	var output []byte
	for _, v := range slc {
		if contentRune(r, rune(v)) {
			continue
		}
		output = append(output, v)
	}
	return output
}

func contentRune(list []rune, r rune) bool {
	var b bool
	for _, v := range list {
		if v == r {
			b = true
			break
		}
	}
	return b
}
