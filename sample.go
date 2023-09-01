package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 175*time.Millisecond)
	go func() {
		defer cancel()
		now := time.Now()
		<-ctx.Done()
		fmt.Println("Timeout:", time.Since(now))
	}()

	time.Sleep(500 * time.Millisecond)
}
