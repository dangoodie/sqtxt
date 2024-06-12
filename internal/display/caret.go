package display

import (
	"image/color"

	"fyne.io/fyne/v2/theme"
)

type CaretStyle struct{}

func (c CaretStyle) TextColor() color.Color {
	return theme.BackgroundColor()
}

func (c CaretStyle) BackgroundColor() color.Color {
	return theme.ForegroundColor()
}
