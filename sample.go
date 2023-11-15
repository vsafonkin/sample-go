package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("---")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "https://2ip.ru", nil)
	if err != nil {
		log.Fatal("create request error:", err)
	}

	req.Header.Add("user-agent", "curl")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("request error:", err)
	}

	buf := make([]byte, 1500)
	n, err := resp.Body.Read(buf)
	if err != nil {
		log.Fatal("read body error:", err)
	}
	fmt.Println(string(buf[:n:n]))
}
