package main

import (
	"fmt"

	// "github.com/vsafonkin/sample-go/db"
	"github.com/vsafonkin/sample-go/journalctl"
)

func main() {

	logs, err := journalctl.Journalctl()
	if err != nil {
		panic(err)
	}
	fmt.Println(logs[0].Hostname)

	// conn, err := db.NewConnect("localhost", "5432", "dvdrental", "postgres", "admin")
	// if err != nil {
	// 	panic(err)
	// }

	for _, v := range logs {
		fmt.Printf("%+v\n", v)
		fmt.Println()
		fmt.Println()
		_ = fmt.Sprintf(`INSERT INTO journal
			(
			log_timestamp,
			log_datetime,
			hostname,
			pid,
			message,
			syslog_identifier,
			exe,
			cmdline,
			system_unit,
			transport,
			priority
			) 
			VALUES ($$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$, $$%s$$);`,
			v.Timestamp,
			v.SyslogDate,
			v.Hostname,
			v.PID,
			v.Message,
			v.SyslogIdentifier,
			v.Exe,
			v.CmdLine,
			v.SystemdUnit,
			v.Transport,
			v.Priority)
		// conn.ExecRawQuery(query)
	}
}
