package main

import (
	"context"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	num := 1
	duration := 1 * time.Second
	query := "SELECT * FROM company;"
	TestDB(num, query, "goapp", duration)
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
	freq := 1 * time.Nanosecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(duration)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
