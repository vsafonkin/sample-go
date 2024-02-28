package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host = "localhost"
var port = "5432"
var user = "postgres"
var password = "admin"
var dbname = "dvdrental"

type Actor struct {
	ActorId   int    `json:"actor_id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func ConnectToDatabase() {
	// dsn := "postgres://postgres:admin@localhost:5432/dvdrental"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("conn:", conn)
	var actor Actor
	conn.First(&actor)
	fmt.Println(actor)
	var actor2 Actor
	conn.First(&actor2, 3)
	fmt.Println(actor2)

	var allActors []Actor
	conn.Find(&allActors)
	fmt.Printf("%v+\n", allActors)
}

func (Actor) TableName() string {
	return "actor"
}
