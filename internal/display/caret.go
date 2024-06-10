package display

import (
	"image/color"

	structs "github.com/dangoodie/sqtxt/internal/structs"
    util "github.com/dangoodie/sqtxt/pkg/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Caret struct {
	*canvas.Rectangle
}

func NewCaret(p structs.Position) *Caret {
	// Get the configuration
	cfg := util.GetConfig()

	c := &Caret{
		Rectangle: canvas.NewRectangle(&color.RGBA{255, 0, 0, 255}), // red for now
	}

	c.Rectangle.Move(convertPositionToScreen(p))
	c.Resize(fyne.Size{Height: float32(cfg.RowHeight), Width: float32(cfg.ColWidth)})

	return c
}

func (c *Caret) Move(p structs.Position) {
	c.Rectangle.Move(convertPositionToScreen(p))
}

func convertPositionToScreen(p structs.Position) fyne.Position {
	// Get the configuration
	cfg := util.GetConfig()

	row := float32(p.Row * cfg.RowHeight)
	col := float32(p.Col * cfg.ColWidth)

	return fyne.NewPos(col, row)
}
