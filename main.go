package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/ericfialkowski/goterm"
)

var columns int
var rows int

var world [][]bool
var neighbors [][]int

func main() {
	columns = goterm.Width()
	rows = goterm.Height() - 1

	world = make([][]bool, rows)
	neighbors = make([][]int, rows)

	rand.Seed(time.Now().UnixNano())
	for x := 0; x < rows; x++ {
		world[x] = make([]bool, columns)
		neighbors[x] = make([]int, columns)

		for y := 0; y < columns; y++ {
			world[x][y] = rand.Intn(10) < 5 // random world
		}
	}

	//world[0][1] = true
	//world[1][1] = true
	//world[2][1] = true

	for gen := 0; gen < 1000; gen++ {
		printWorld(gen)
		changed := nextGen()
		if !changed {
			goterm.Println(strings.Repeat("\n", 2))
			goterm.Println("stale world")
			goterm.Flush()
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func printWorld(gen int) {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	//goterm.Printf("Generation: %d\t%d\t%d\n", gen, goterm.Width(), goterm.Height())
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if world[x][y] {
				goterm.Print("#")
			} else {
				goterm.Print(" ")
			}
		}
		goterm.Println()
	}
	goterm.Flush()
}

func nextGen() bool {
	worldChanged := false

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			countNeighbors(x, y)
		}
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			alive, changed := liveOrDie(x, y)
			world[x][y] = alive
			if changed {
				worldChanged = true
			}
		}
	}
	return worldChanged
}

func countNeighbors(x, y int) {
	neighbors[x][y] = 0

	// count neighbors
	for xOffset := -1; xOffset < 2; xOffset++ {
		for yOffset := -1; yOffset < 2; yOffset++ {
			newX := x + xOffset
			newY := y + yOffset
			//fmt.Printf("\t(%d,%d)\n", newX, newY)
			if newX > -1 && newY > -1 && newX < rows && newY < columns {
				if world[newX][newY] {
					neighbors[x][y]++
				}
			}
		}
	}
	if world[x][y] {
		neighbors[x][y]-- // counted self
	}

	//fmt.Printf("(%d,%d) = %d\n", x, y, neighbors)
}

func liveOrDie(x, y int) (bool, bool) {
	alive := false
	n := neighbors[x][y]

	if world[x][y] {
		alive = n == 2 || n == 3
	} else {
		alive = n == 3
	}

	changed := world[x][y] != alive

	return alive, changed
}
