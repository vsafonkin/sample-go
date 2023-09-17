package progressbar

import (
	"fmt"
	"math/rand"
	"time"
)

func ProgressBar(done chan<- struct{}) {
	var n int
	mark := "*"
	for n <= 100 {
		fmt.Printf("\r[%d%%]%s", n, mark)
		time.Sleep(RandTimeout(300))
		if n%4 == 0 {
			mark += "*"
		}
		n++
	}
	fmt.Println()
	done <- struct{}{}
}

func RandTimeout(limit int) time.Duration {
	var timeout time.Duration
	randInt := rand.Intn(limit)
	timeout = time.Duration(randInt)
	return timeout * time.Millisecond
}
