package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	num := 1
	duration := 10 * time.Second
	var wg sync.WaitGroup
	for _ = range 5 {
		wg.Add(1)
		query := fmt.Sprintf("SELECT * FROM airports_data;")
		go func() {
			defer wg.Done()
			TestDB(num, query, "goapp", duration)
		}()
	}
	wg.Wait()
}

func TestDB(numConn int, query, appname string, duration time.Duration) {
	config := db.Config{
		Host:    "localhost",
		Port:    "5432",
		DBName:  "demo",
		User:    "postgres",
		Pass:    "admin",
		AppName: appname,
	}

	pool, err := db.NewPool(numConn, config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	freq := 0 * time.Nanosecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(duration)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
