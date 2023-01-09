package world

const (
	ALIVE = 1
	FREE  = 0
)

// World represent the current world in game of life
type World struct {
	matrix [][]int
	Size   int
}

// Generator function type that give a generator or random number, in future could have
// the x and y as parameter, so can have logic based on position
type Generator func() bool

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
		Size:   size,
		matrix: slice,
	}
}

func (w *World) countNeighbour(x, y int) int {
	neightbour := 0
	for i := x-1; i < x+2; i++ {
		for j := y-1; j < y+2; j++ {
			if i == x && j == y {
				continue
			}
			neightbour += w.valueAt(i, j)
		}
	}
	return neightbour
}

func (w *World) valueAt(x, y int) int {
    var normX = (x + w.Size) % w.Size
    var normY = (y + w.Size) % w.Size
    return w.matrix[normY][normX]
}
func (w *World) canSurvive(x, y int) bool {
	neightbour := w.countNeighbour(x, y)
	return neightbour == 2 || neightbour == 3
}

func (w *World) willBorn(x, y int) bool {
	neightbour := w.countNeighbour(x, y)
	return neightbour == 3
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
	newWorld := NewWorld(w.Size)
	for x := 0; x < w.Size; x++ {
		for y := 0; y < w.Size; y++ {
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

func Random(size int, fn Generator) *World {
	w := NewWorld(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			alive := fn()
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
