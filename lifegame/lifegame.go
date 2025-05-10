package lifegame

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 40
	height = 20
)

// Создание сетки
func createGrid() [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

// Инициализация случайных активных клеток
func initRandom(grid [][]bool) {
	rand.Seed(time.Now().UnixNano())
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[y][x] = rand.Intn(2) == 0 // 50% шанс активной клетки
		}
	}
}

// Подсчет соседей
func countNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue // Пропуск текущей клетки
			}
			ny := y + dy
			nx := x + dx
			if ny >= 0 && ny < height && nx >= 0 && nx < width {
				if grid[ny][nx] {
					count++
				}
			}
		}
	}
	return count
}

// Обновление состояния клеток
func updateGrid(grid [][]bool) [][]bool {
	newGrid := createGrid()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			neighbors := countNeighbors(grid, x, y)
			// Правила игры «Жизнь»
			if grid[y][x] {
				newGrid[y][x] = neighbors == 2 || neighbors == 3
			} else {
				newGrid[y][x] = neighbors == 3
			}
		}
	}
	return newGrid
}

// Отрисовка сетки в терминале
func drawGrid(grid [][]bool) {
	fmt.Print("\033[H\033[2J") // Очистка терминала
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] {
				fmt.Print("*") // Активная клетка
			} else {
				fmt.Print(" ") // Пустое место
			}
		}
		fmt.Println()
	}
	time.Sleep(100 * time.Millisecond) // Задержка между кадрами
}

func Run() {
	grid := createGrid()
	initRandom(grid)

	// Бесконечный цикл симуляции
	for {
		drawGrid(grid)
		grid = updateGrid(grid)
	}
}
