package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadSyslog(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var counter int
	var out []string
	var temp string
	var e error
	for {
		temp, e = reader.ReadString('\n')
		if e != nil {
			break
		}
		out = append(out, temp)
		counter++
	}
	fmt.Println("Read lines from syslog:", counter)
	return out, nil
}

func Parse(s string) []string {
	temp := strings.Split(s, " ")
	month := temp[0]
	logdate := temp[2]
	logtime := temp[3]
	hostname := temp[4]
	proc := temp[5]
	body := strings.Join(temp[6:], " ")
	return []string{month, logdate, logtime, hostname, proc, body[:len(body)-1]}
}
