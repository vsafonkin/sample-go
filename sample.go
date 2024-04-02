package main

import (
	"context"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	TestDB(1, 100*time.Second)
}

func TestDB(numConn int, timeout time.Duration) {
	config := db.Config{
		Host:    "localhost",
		Port:    "5432",
		DBName:  "dvdrental",
		User:    "postgres",
		Pass:    "admin",
		AppName: "goapp",
	}

	pool, err := db.NewPool(numConn, config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	query := "SELECT * FROM person;"
	freq := 1 * time.Microsecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(timeout)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
