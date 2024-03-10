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

	var td time.Duration
	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// randomUpdate(conn)
				td = selectKernel(conn)
				counter++
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	go func() {
		for {
			fmt.Printf("counter: %d, request time: %v\r", counter, td)
			time.Sleep(100 * time.Millisecond)
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

func selectKernel(conn *db.DB) time.Duration {
	query := fmt.Sprintf(
		`SELECT exe, message FROM journal WHERE message ~ '%s';`,
		"USER=root",
	)
	return conn.ExecRawQuery(query)
}
