package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/sciack/life/world"
)

const (
	SIZE = 15
)

func initScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	return s
}

func printText(w *world.World, s tcell.Screen, blank, alive tcell.Style) {
	drawBox(s, 0, 0, SIZE+2, SIZE+2, alive)
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			if w.IsAlive(x, y) {
				s.SetContent(x+1, y+1, tcell.RuneCkBoard, nil, alive)
			} else {
				s.SetContent(x+1, y+1, ' ', nil, alive)
			}
		}
	}
	s.Show()

}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

}

func main() {
	s := initScreen()

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	alive := tcell.StyleDefault.Foreground(tcell.ColorLime).Background(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	w := world.Random(SIZE)
	for i := 0; i < 100; i++ {
		printText(w, s, defStyle, alive)
		time.Sleep(100 * time.Millisecond)
		w = w.Next()
	}
	s.Clear()
}
