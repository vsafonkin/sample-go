package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewConnect(host, port, dbname, user, password string) (*DB, error) {
	// dsn := "postgres://postgres:admin@localhost:5432/dvdrental"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

func (db *DB) ExecRawQuery(query string) {
	db.conn.Raw(query).Scan(&struct{}{})
}
