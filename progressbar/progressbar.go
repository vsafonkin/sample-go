package progressbar

import (
	"fmt"
	"time"
)

func ProgressBar(done chan<- struct{}) {
	var n int
	mark := "*"
	for n <= 100 {
		fmt.Printf("\r[%d%%]%s", n, mark)
		time.Sleep(100 * time.Millisecond)
		if n%4 == 0 {
			mark += "*"
		}
		n++
	}
	fmt.Printf("\n")
	done <- struct{}{}
}
