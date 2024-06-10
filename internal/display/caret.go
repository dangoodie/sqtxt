package display

import (
	"image/color"

	structs "github.com/dangoodie/sqtxt/internal/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Caret struct {
	*canvas.Rectangle
}

// ROW_HEIGHT and COL_WIDTH are the dimensions of the caret
const (
	ROW_HEIGHT = 20
	COL_WIDTH  = 10
)

func NewCaret(p structs.Position) *Caret {
	c := &Caret{
		Rectangle: canvas.NewRectangle(&color.RGBA{255, 0, 0, 255}), // red for now
	}

	c.Rectangle.Move(convertPositionToScreen(p))
	c.Resize(fyne.Size{Height: float32(ROW_HEIGHT), Width: float32(COL_WIDTH)})

	return c
}

func (c *Caret) Move(p structs.Position) {
	c.Rectangle.Move(convertPositionToScreen(p))
}

func convertPositionToScreen(p structs.Position) fyne.Position {
	row := float32(p.Row * ROW_HEIGHT)
	col := float32(p.Col * COL_WIDTH)

	return fyne.NewPos(row, col)
}
