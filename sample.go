package main

import (
	"fmt"
	"os"
)

type TestError struct {
	Test      string
	Message   string
	ErrStatus int
}

func (e TestError) Error() string {
	return fmt.Sprintf("error message: %s, error status: %d\n", e.Message, e.ErrStatus)
}

func main() {
	fmt.Println("-------------")
	if _, err := openFile("./makefilee"); err != nil {
		fmt.Println(err)
	}

	recPanic()
	fmt.Println("goodbye")
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func throwPanic() {
	panic(TestError{
		Test:      "test message",
		Message:   "stop program",
		ErrStatus: 123,
	})
}

func recPanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("recover panic: %v\n", v)
		}
	}()
	throwPanic()
}
