package display

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	editor "github.com/dangoodie/sqtxt/internal/editor"
)

type Screen struct {
	editor    *editor.Editor
	container *fyne.Container
	textGrid  *widget.TextGrid
}

func NewScreen(e *editor.Editor) *Screen {
	s := &Screen{
		editor:    e,
		container: container.NewVBox(),
        textGrid: widget.NewTextGridFromString(e.Buffer.Read()),
	}

	s.textGrid.ShowLineNumbers = true
	s.textGrid.SetStyle(e.Cursor.Row, e.Cursor.Col, CaretStyle{})
	s.container.Add(s.textGrid)

	return s
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
	window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) { // This is the event handler for key presses
		e.HandleKeyInput(string(ev.Name))
        screen.textGrid.SetText(e.Buffer.Read())
		screen.container.Refresh()
	})

	window.Canvas().SetOnTypedRune(func(r rune) { // This is the event handler for typing characters
		e.HandleRuneInput(r)
        screen.textGrid.SetText(e.Buffer.Read())
		screen.container.Refresh()
	})

	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
