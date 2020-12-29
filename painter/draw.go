package painter

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// Painter struct, contain the reference to the screen and other useful information
type Painter struct {
	screen   tcell.Screen
	alive    tcell.Style
	defStyle tcell.Style
	vOffset  int
}

//New create a new Painter struct containing the screen and other useful information
func New() *Painter {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	alive := tcell.StyleDefault.Foreground(tcell.ColorLime).Background(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	return &Painter{screen: s, defStyle: defStyle, alive: alive}
}

func (d *Painter) drawBox(x1, y1, x2, y2 int) {
	s := d.screen
	style := d.defStyle
	x2 = x2 * 2
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

// Clear clear the screen
func (d *Painter) Clear() {
	d.screen.Clear()
}

// DrawAlive draw an alive cell
func (d *Painter) DrawAlive(x, y int) {
	d.drawSymbol(x, y, tcell.RuneCkBoard, d.alive)
}

func (d *Painter) drawSymbol(x, y int, symbol rune, style tcell.Style) {
	xnorm := x*2 + 1 // zero is the border
	d.screen.SetContent(xnorm, y+1+d.vOffset, symbol, nil, style)
	d.screen.SetContent(xnorm+1, y+1+d.vOffset, symbol, nil, style)
}

//DrawEmpty draw an empty cell
func (d *Painter) DrawEmpty(x, y int) {
	d.drawSymbol(x, y, ' ', d.defStyle)
}

//StartDrawing start an empty canvass read for the drawing
func (d *Painter) StartDrawing(size, vOffset int) {
	d.vOffset = vOffset
	d.drawBox(0, vOffset, size, size+vOffset)
}

// EndDrawing show the changes
func (d *Painter) EndDrawing() {
	d.screen.Show()
}

// DrawText draw a text starting from the x,y coordinates using the default color
func (d *Painter) DrawText(x, y int, text string) {
	for index, r := range []rune(text) {
		d.screen.SetContent(x+index, y, r, nil, d.defStyle)
	}
}

// DrawTextHigh dsame as DrawText but with the same color as Alive
func (d *Painter) DrawTextHigh(x, y int, text string) {
	for index, r := range []rune(text) {
		d.screen.SetContent(x+index, y, r, nil, d.alive)
	}
}
