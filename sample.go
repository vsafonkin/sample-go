package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/vsafonkin/sample-go/db"
	// "github.com/vsafonkin/sample-go/journalctl"
)

func main() {

	// logs, err := journalctl.Parse()
	// if err != nil {
	// panic(err)
	// }
	// fmt.Println(logs[0].Hostname)

	conn, err := db.NewConnect("localhost", "5432", "dvdrental", "postgres", "admin")
	if err != nil {
		panic(err)
	}

	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// randomUpdate(conn)
				selectKernel(conn)
				counter++
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}
	go func() {
		for {
			fmt.Printf("counter: %v\r", counter)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}

func randomUpdate(conn *db.DB) {
	query := fmt.Sprintf(
		`UPDATE person SET age = %d WHERE id = %d`,
		rand.Intn(60)+1, rand.Intn(9)+1)
	conn.ExecRawQuery(query)
}

func selectKernel(conn *db.DB) {
	query := fmt.Sprintf(
		`SELECT * FROM journal WHERE message ~ '%s' ORDER BY log_timestamp`,
		"USER=root",
	)
	conn.ExecRawQuery(query)
}
