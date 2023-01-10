package main

import (
    "math/rand"
    "strconv"
    "time"

    "github.com/sciack/life/painter"
    "github.com/sciack/life/world"
)

const (
	// THRESHOLD tune the random grid population, higher the number less populated is the grid (<100)
	THRESHOLD = 80
)

func drawHeader(paint *painter.Painter, iteration int) {
	paint.DrawText(0, 2, "Iteration")
	paint.DrawTextHigh(len("Iteration")+1, 2, strconv.Itoa(iteration))

}
func drawWorld(w *world.World, paint *painter.Painter) {
	var size = w.Size
	paint.StartDrawing(size+1, 4)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if w.IsAlive(x, y) {
				paint.DrawAlive(x, y)
			} else {
				paint.DrawEmpty(x, y)
			}
		}
	}

	paint.EndDrawing()
}

func randomGenerator() bool {
    return rand.Intn(100) >= THRESHOLD
}

func emit(worlds chan *world.World, size int) {
	w := world.Random(size, randomGenerator)
	worlds <- w
	for w.CountAlive() > 0 {
		worlds <- w
		w = w.Next()
	}
	close(worlds)
}

func main() {
	paint := painter.New()
	var width, height = paint.GridSize()

	var size = Min(width, height)

	rand.Seed(time.Now().UnixNano())

	var worlds = make(chan *world.World, 5)
	var i = 0
	go emit(worlds, size)
	for w := range worlds {
		drawHeader(paint, i)
		drawWorld(w, paint)
        time.Sleep(100 * time.Millisecond)
		paint.Interrupted()
		i += 1
	}

	paint.EndDrawing()
	paint.EndSession()
}


func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}