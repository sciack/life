package main

import (
	"fmt"
	"time"

	"github.com/sciack/life/painter"
	"github.com/sciack/life/world"
)

const (
	// SIZE is the size of the grid
	SIZE = 15
	// THRESHOLD tune the random grid population, higher the number less populated is the grid (<100)
	THRESHOLD = 98
)

func printText(w *world.World, paint *painter.Painter, iteration int) {
	paint.DrawText(0, 0, fmt.Sprintf("Iteration %v", iteration))
	paint.StartDrawing(SIZE+1, 2)
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

	w := world.Random(SIZE, THRESHOLD)
	for i := 0; i < 100 && w.CountAlive() > 0; i++ {
		printText(w, paint, i)
		time.Sleep(200 * time.Millisecond)
		w = w.Next()
	}

}
