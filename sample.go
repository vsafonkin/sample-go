package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows    = 30
	columns = 50
)

type grid [rows][columns]bool

func main() {
	var g grid
	rand.Seed(time.Now().UnixNano())
	randomize(&g)
	draw(g)
	for {
		evolve(&g)
		draw(g)
		time.Sleep(200 * time.Millisecond)
	}
}

func randomize(g *grid) {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			g[i][j] = rand.Intn(10) == 0
		}
	}
}

func draw(g grid) {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if g[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func evolve(g *grid) {
	var ng grid
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			n := neighbors(g, i, j)
			if g[i][j] && (n == 2 || n == 3) {
				ng[i][j] = true
			} else if !g[i][j] && n == 3 {
				ng[i][j] = true
			}
		}
	}
	*g = ng
}

func neighbors(g *grid, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= rows {
				continue
			}
			if y+j < 0 || y+j >= columns {
				continue
			}
			if g[x+i][y+j] {
				count++
			}
		}
	}
	return count
}
