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
	SIZE = 15
	// THRESHOLD tune the random grid population, higher the number less populated is the grid (<100)
	THRESHOLD = 97
)

func drawHeader(paint *painter.Painter, iteration int) {
	paint.DrawText(0, 2, "Iteration")
	paint.DrawTextHigh(len("Iteration")+1, 2, strconv.Itoa(iteration))

}
func drawWorld(w *world.World, paint *painter.Painter) {
	paint.StartDrawing(SIZE+1, 4)

	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
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
	rand.Seed(time.Now().UnixNano())

	generator := func() bool {
		return rand.Intn(100) >= THRESHOLD
	}

	w := world.Random(SIZE, generator)
	for i := 1; i <= 100 && w.CountAlive() > 0; i++ {
		drawHeader(paint, i)
		drawWorld(w, paint)
		time.Sleep(200 * time.Millisecond)
		w = w.Next()
	}

	paint.EndDrawing()
}
