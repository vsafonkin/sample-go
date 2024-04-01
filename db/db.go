package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

type Config struct {
	Host, Port, DBName, User, Pass, AppName string
}

func NewConnect(config Config) (*DB, error) {
	// dsn := "postgres://postgres:admin@localhost:5432/dvdrental"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s application_name=%s",
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.DBName,
		config.AppName)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

func (db *DB) ExecRawQuery(query string) time.Duration {
	start := time.Now()
	db.conn.Raw(query).Scan(&struct{}{})
	return time.Since(start)
}

func NewPool(n int, config Config) ([]*DB, error) {
	var connPool []*DB
	for i := range n {
		config.AppName = fmt.Sprintf("goapp_%d", i)
		conn, err := NewConnect(config)
		if err != nil {
			fmt.Println("appname:", config.AppName, "connection error:", err)
			continue
		}
		connPool = append(connPool, conn)
	}
	if len(connPool) == 0 {
		return nil, fmt.Errorf("there are no any successful connections, pool is empty")
	}
	return connPool, nil
}

func TestLoad(ctx context.Context, query string, pool []*DB, freq time.Duration) {
	num := len(pool)
	var dur time.Duration
	var counter int
	go func() {
	out:
		for {
			fmt.Printf("counter: %v, time: %v    \r", counter, dur)
			select {
			case <-ctx.Done():
				fmt.Println()
				break out
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	for i := range num {
		go func() {
			conn := pool[i]
			defer func() {
				sql, err := conn.conn.DB()
				if err != nil {
					fmt.Println("close connection error:", err)
					return
				}
				sql.Close()
			}()
		out:
			for {
				dur = conn.ExecRawQuery(query)
				select {
				case <-ctx.Done():
					break out
				default:
					time.Sleep(freq)
					counter++
				}
			}
		}()
	}
	select {
	case <-ctx.Done():
	}
}
