package main

import (
	"context"
	"time"

	"github.com/vsafonkin/sample-go/db"
)

func main() {
	config := db.Config{
		Host:    "localhost",
		Port:    "5432",
		DBName:  "dvdrental",
		User:    "postgres",
		Pass:    "admin",
		AppName: "goapp",
	}

	pool, err := db.NewPool(8, config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	query := "SELECT * FROM person;"
	freq := 100 * time.Millisecond
	go db.TestLoad(ctx, query, pool, freq)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(300 * time.Millisecond)
}
