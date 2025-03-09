package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func main() {
	wordHash := "8b215d105e88b6e45c4c6af99f098398"

	ch := make(chan struct{})
	for i := 0; i < 16; i++ {
		go func() {
			for {
				s, hash := genMD5RandHashString(7)
				if hash == wordHash {
					println(s)
					ch <- struct{}{}
					break
				}
			}
		}()
	}
	<-ch
}

func genMD5RandHashString(n int) (string, string) {
	symbols := "qwertyuiopasdfghjklzxcvbnm"
	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}

	mdHash := md5.Sum(buf)
	hashString := hex.EncodeToString(mdHash[:])
	return string(buf), hashString
}
