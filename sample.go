package main

import (
	"fmt"
	"os"

	"github.com/vsafonkin/sample-go/db"
	"github.com/vsafonkin/sample-go/parser"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}

	conn, err := db.NewConnect("localhost", "5432", "dvdrental", "postgres", "admin")
	if err != nil {
		panic(err)
	}

	logs, err := parser.ReadSyslog("/var/log/syslog")
	if err != nil {
		panic(err)
	}

	for _, v := range logs {
		s := parser.Parse(v)
		query := fmt.Sprintf("INSERT INTO syslog (month, logdate, logtime, hostname, process, message) VALUES ($$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$);", s[0], s[1], s[2], s[3], s[4], s[5])
		conn.ExecRawQuery(query)
	}
}
