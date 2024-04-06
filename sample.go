package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	queries := []string{
		"SELECT * FROM airports_data;",
		"SELECT * FROM aircrafts_data;",
		"SELECT * FROM seats;",
	}
	num := 1
	duration := 600 * time.Second
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		query := queries[i]
		go func() {
			defer wg.Done()
			appname := fmt.Sprintf("goapp_%d", i)
			TestDB(num, query, appname, duration)
		}()
	}
	wg.Wait()
}

func TestDB(numConn int, query, appname string, duration time.Duration) {
	config := db.Config{
		Host:    "192.168.18.131",
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
	freq := 100 * time.Millisecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(duration)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
