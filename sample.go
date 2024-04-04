package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	num := 1
	duration := 1 * time.Second
	for _ = range 64 {
		query := fmt.Sprintf("UPDATE company SET stock_price = %d WHERE id = 1;", rand.Intn(5000))
		TestDB(num, query, "goapp", duration)
		time.Sleep(10 * time.Millisecond)
	}
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
