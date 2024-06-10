package display

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	editor "github.com/dangoodie/sqtxt/internal/editor"
)

type Screen struct {
	editor    *editor.Editor
	caret     *Caret
	textObj   []*canvas.Text
	container *fyne.Container
}

func NewScreen(e *editor.Editor) *Screen {
	s := &Screen{
		editor: e,
		caret:  NewCaret(e.Cursor),
	}
	s.loadText()
	cObj := []fyne.CanvasObject{}
	for _, text := range s.textObj {
		cObj = append(cObj, text)
	}
	cObj = append(cObj, s.caret.Rectangle)
	s.container = container.NewWithoutLayout(cObj...)

	return s
}

func (s *Screen) loadText() {
	s.textObj = []*canvas.Text{}
	for i, line := range s.editor.Buffer.Read() {
		text := canvas.NewText(string(line), &color.RGBA{255, 255, 255, 255})
		text.Move(fyne.NewPos(float32(i*ROW_HEIGHT), 0))
		s.textObj = append(s.textObj, text)
	}
}

// InitializeScreen sets up the terminal screen
func Start(e editor.Editor) {
	// Code to initialize the screen
	fmt.Println("Initializing screen...")
	app := app.New()
	title := "sqtxt - " + e.Filename
	window := app.NewWindow(title)

	screen := NewScreen(&e)

	window.SetContent(screen.container)
	window.Canvas().SetOnTypedRune(func(r rune) {
		e.HandleInput(r)
		screen.loadText()
		screen.caret.Move(e.Cursor)
		screen.container.Refresh()
	})

	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
