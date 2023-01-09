package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/sciack/life/painter"
	"github.com/sciack/life/world"
)

const (
	// SIZE is the size of the grid
	SIZE = 20
	// THRESHOLD tune the random grid population, higher the number less populated is the grid (<100)
	THRESHOLD = 80

    ITERATION = 1000
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

func main() {
	paint := painter.New()
	var width, height = paint.ScreenSize()

	var size = 0
	if width < height {
		size = width
	} else {
		size = height
	}

	rand.Seed(time.Now().UnixNano())

	generator := func() bool {
		return rand.Intn(100) >= THRESHOLD
	}

	w := world.Random(size-6, generator)
	for i := 1; i <= ITERATION && w.CountAlive() > 0; i++ {

		drawHeader(paint, i)
		drawWorld(w, paint)
		time.Sleep(200 * time.Millisecond)
        paint.Interrupted()
		w = w.Next()
	}

	paint.EndDrawing()
}
