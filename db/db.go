package db

import (
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
	fmt.Printf("New connect to db %s:%s\n", config.Host, config.Port)
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

func (db *DB) ExecRawQuery(query string) (any, time.Duration) {
	start := time.Now()
	out := db.conn.Raw(query).Row()
	return out, time.Since(start)
}

func (db *DB) Row(tablename string, column string, value string) map[string]any {
	out := make(map[string]any, 1)
	db.conn.Table(tablename).Take(out, column, value)
	return out
}
