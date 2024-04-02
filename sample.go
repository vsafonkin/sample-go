package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	num := 1
	duration := 10 * time.Second
	var wg sync.WaitGroup
	for i := range 32 {
		wg.Add(1)
		query := fmt.Sprintf("UPDATE person SET age = %d WHERE id = %d;", rand.Intn(60)+1, rand.Intn(12)+1)
		go func() {
			defer wg.Done()
			appname := fmt.Sprintf("goapp_%d", i)
			TestDB(num, query, appname, duration)
		}()
	}
	wg.Wait()
}

func TestDB(numConn int, query, appname string, timeout time.Duration) {
	config := db.Config{
		Host:    "localhost",
		Port:    "5432",
		DBName:  "dvdrental",
		User:    "postgres",
		Pass:    "admin",
		AppName: appname,
	}

	pool, err := db.NewPool(numConn, config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	freq := 1 * time.Nanosecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(timeout)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
