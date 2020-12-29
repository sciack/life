package main

import (
	"fmt"
	"math/rand"
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

func printText(w *world.World, paint *painter.Painter, iteration int) {
	paint.DrawText(0, 2, fmt.Sprintf("Iteration %v", iteration))
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
	for i := 0; i < 100 && w.CountAlive() > 0; i++ {
		printText(w, paint, i)
		time.Sleep(200 * time.Millisecond)
		w = w.Next()
	}

}
