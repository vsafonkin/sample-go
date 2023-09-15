package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	f, err := os.Open("./test_data/users.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

	buf := make([]byte, 128)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
