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
	if err := wrap("./makefilee"); err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("goodbye")
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func throwError(path string) error {
	if _, err := openFile(path); err != nil {
		return fmt.Errorf("throwError func: %w", err)
	}
	return nil
}

func wrap(path string) error {
	return fmt.Errorf("wrap func: %w", throwError(path))
}
