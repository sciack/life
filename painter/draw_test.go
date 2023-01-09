package painter

import (
    "log"
    "testing"

	"github.com/gdamore/tcell/v2"
)

func painterForTesting() (*Painter, tcell.SimulationScreen) {
	screen := tcell.NewSimulationScreen("utf-8")
    if err := screen.Init(); err != nil {
        log.Fatalf("%+v", err)
    }
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	alive := tcell.StyleDefault.Foreground(tcell.ColorLime).Background(tcell.ColorReset)
	return &Painter{
		screen:   screen,
		alive:    alive,
		defStyle: defStyle,
		vOffset:  0,
		hOffset:  0,
	}, screen
}

func TestDrawAlive(t *testing.T) {
	p, screen := painterForTesting()
	p.DrawAlive(0, 0)
	p.EndDrawing()
	cells, width, _ := screen.GetContents()

	got := cells[1*width+1]

	if got.Runes[0] != tcell.RuneCkBoard {
		t.Errorf("Expected %v got %v", tcell.RuneCkBoard, got.Runes[0])
	}

	got = cells[1*width+2]
	if got.Runes[0] != tcell.RuneCkBoard {
		t.Errorf("Expected %v got %v", tcell.RuneCkBoard, got.Runes[0])
	}

}

func TestDrawEmpty(t *testing.T) {
	p, screen := painterForTesting()
	p.DrawEmpty(0, 0)
	p.EndDrawing()
	cells, width, _ := screen.GetContents()

	got := cells[1*width+1]

	if got.Runes[0] != ' ' {
		t.Errorf("Expected %q got %q", ' ', got.Runes[0])
	}

	got = cells[1*width+2]
	if got.Runes[0] != ' ' {
		t.Errorf("Expected %q got %q", ' ', got.Runes[0])
	}

}

func TestDrawText(t *testing.T) {
	p, screen := painterForTesting()
	expected := "Text"
	p.DrawText(0, 0, expected)
	p.EndDrawing()
	cells, _, _ := screen.GetContents()
	style := cells[0].Style
	got := cellsToString(cells[:len(expected)])

	if got != expected {
		t.Errorf("Expected %q got %q", expected, got)
	}
	if style != p.defStyle {
		t.Errorf("Expected style %q, got %q", p.defStyle, style)
	}
}

func cellsToString(cells []tcell.SimCell) string {
	var runes []rune
	for _, cell := range cells {
		runes = append(runes, cell.Runes[0])
	}
	return string(runes)
}

func TestDrawTextHigh(t *testing.T) {
	p, screen := painterForTesting()
	expected := "Text"
	p.DrawTextHigh(0, 0, expected)
	p.EndDrawing()
	cells, _, _ := screen.GetContents()
	style := cells[0].Style
	got := cellsToString(cells[:len(expected)])

	if got != expected {
		t.Errorf("Expected %q got %q", expected, got)
	}
	if style != p.alive {
		t.Errorf("Expected style %q, got %q", p.defStyle, style)
	}
}
