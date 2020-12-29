package world

import (
	"math/rand"
	"time"
)

const (
	ALIVE = 1
	FREE  = 0
)

// World represent the current world in game of life
type World struct {
	matrix [][]int
	size   int
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// NewWorld create a world struct with the predefined size
func NewWorld(size int) *World {
	slice := [][]int{}
	for i := 0; i < size; i++ {
		line := []int{}
		for k := 0; k < size; k++ {
			line = append(line, 0)
		}
		slice = append(slice, line)
	}
	return &World{
		size:   size,
		matrix: slice,
	}
}

func (w *World) countNeighbour(x, y int) int {
	neightbour := 0
	for i := max(x-1, 0); i < min(x+2, w.size); i++ {
		for k := max(y-1, 0); k < min(y+2, w.size); k++ {
			if i == x && k == y {
				continue
			}
			neightbour += w.matrix[k][i]
		}
	}
	return neightbour
}

func (w *World) canSurvive(x, y int) bool {
	neightbour := w.countNeighbour(x, y)
	return neightbour == 2 || neightbour == 3
}

func (w *World) willBorn(x, y int) bool {
	neightbour := w.countNeighbour(x, y)
	return neightbour == 2
}

func (w *World) SetState(x, y int, alive bool) {
	if alive {
		w.matrix[y][x] = ALIVE
	} else {
		w.matrix[y][x] = FREE
	}
}

func (w *World) IsAlive(x, y int) bool {
	return w.matrix[y][x] == ALIVE
}

// Next will give the next world configuration
func (w *World) Next() *World {
	newWorld := NewWorld(w.size)
	for x := 0; x < w.size; x++ {
		for y := 0; y < w.size; y++ {
			if w.canSurvive(x, y) && w.IsAlive(x, y) {
				newWorld.SetState(x, y, w.IsAlive(x, y))
			} else if w.willBorn(x, y) {
				newWorld.SetState(x, y, true)
			} else {
				newWorld.SetState(x, y, false)
			}
		}
	}
	return newWorld
}

func Random(size, threshold int) *World {
	rand.Seed(time.Now().UnixNano())
	w := NewWorld(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			alive := rand.Intn(100) >= threshold
			w.SetState(x, y, alive)
		}
	}
	return w
}

func (w *World) CountAlive() int {
	alive := 0
	for _, row := range w.matrix {
		for _, value := range row {
			alive += value
		}

	}
	return alive
}
