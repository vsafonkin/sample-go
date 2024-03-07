package journalctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Jlog struct {
	Hostname         string `json:"_HOSTNAME"`
	PID              string `json:"_PID"`
	SyslogIdentifier string `json:"SYSLOG_IDENTIFIER"`
	Timestamp        string `json:"__REALTIME_TIMESTAMP"`
	Exe              string `json:"_EXE"`
	SyslogDate       string `json:"SYSLOG_TIMESTAMP"`
	CmdLine          string `json:"_CMDLINE"`
	Transport        string `json:"_TRANSPORT"`
	SystemdUnit      string `json:"_SYSTEMD_UNIT"`
	Priority         string `json:"PRIORITY"`
	Message          string `json:"MESSAGE"`
}

func Journalctl() ([]Jlog, error) {
	cmd := exec.Command("journalctl", "-b", "-o", "json")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	logs := strings.Split(string(stdout), "\n")

	var jlog Jlog
	var output []Jlog
	for _, v := range logs {
		if v == "" {
			continue
		}
		err := json.Unmarshal([]byte(v), &jlog)
		if err != nil {
			fmt.Println("parsing error:", err, "string:", v)
			continue
		}
		output = append(output, jlog)
	}
	return output, nil
}
