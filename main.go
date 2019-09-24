package main

import (
	"math/rand"
	"time"

	"github.com/buger/goterm"
)

var columns int
var rows int

var world [][]bool
var prev [][]bool
var prevPrev [][]bool
var neighbors [][]int

func main() {
	columns = goterm.Width()
	rows = goterm.Height() - 1

	world = make([][]bool, rows)
	//prev = make([][]bool,rows)
	neighbors = make([][]int, rows)

	rand.Seed(time.Now().UnixNano())

	for x := 0; x < rows; x++ {
		world[x] = make([]bool, columns)
		//prev[x] = make([]bool, columns)
		neighbors[x] = make([]int, columns)
	}

	for {
		for x := 0; x < rows; x++ {
			for y := 0; y < columns; y++ {
				world[x][y] = rand.Intn(10) < 5 // random world
			}
		}

		for gen := 0; gen < 500; gen++ {
			prevPrev = prev
			prev = world
			printWorld(gen)
			changed := nextGen()
			if !changed {
				break
			}
			//if gen > 2 && genSame(world,prevPrev) {
			//	break
			//}

			time.Sleep(time.Millisecond * 100)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func genSame(current [][]bool, prev [][]bool) bool {
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if current[x][y] != prev[x][y] {
				return false
			}
		}
	}
	return true
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
