package main

import (
	"fmt"
	"math/rand"
    "time"
)

type Game struct {
	width, height int
	grid          [][]bool
}

func NewGame(width, height int) *Game {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return &Game{width: width, height: height, grid: grid}
}

func main() {
	game := NewGame(50, 50)
	game.Init()

		for {
				game.Render()
				game.Update()
				time.Sleep(200 * time.Millisecond)
		}

}

func (g *Game) Init() {
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			g.grid[x][y] = rand.Float64() < 0.25
		}
	}
}

func (g *Game) countAlive(x, y int) int {
		count := 0
		for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
								continue
						}
						nx, ny := x+dx, y+dy
						if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height && g.grid[nx][ny] {
								count++
						}
				}
		}
		return count
} 
func (g *Game) Render() {
	fmt.Print("\033[H\033[2J") // Clear screen
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if g.grid[x][y] {
				fmt.Print("■ ")

			} else {
				fmt.Print("□ ")
			}

		}
		fmt.Println()
	}
}

func (g *Game) Update() {
		newGrid := make([][]bool, g.height)
		for i := range newGrid {
				newGrid[i] = make([]bool, g.width)
		}

		for x := 0; x < g.width; x++ {
				for y := 0; y < g.height; y++ {
						alive := g.countAlive(x, y)
						if  g.grid[x][y] {
								newGrid[x][y] = alive == 2 || alive == 3
						} else {
								newGrid[x][y] = alive == 3
						}
				}
		}

		g.grid = newGrid
}
