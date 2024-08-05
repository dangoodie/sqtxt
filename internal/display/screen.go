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

    // Set the caret style
    x, y := e.Cursor.GetPosition()
	s.textGrid.SetStyle(x, y, CaretStyle{})
	s.container.Add(s.textGrid)

	return s
}

// InitializeScreen sets up the terminal screen
func Start(e editor.Editor) {
	// Code to initialize the screen
	fmt.Println("Initializing screen...")
	app := app.New()
	title := "sqtxt - " + e.Filename
	w := app.NewWindow(title)

	s := NewScreen(&e)

	w.SetContent(s.container)
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) { // This is the event handler for key presses
		e.HandleKeyInput(string(ev.Name))
        s.textGrid.SetText(e.Buffer.Read())
        s.UpdateCaret()
		s.container.Refresh()
	})

	w.Canvas().SetOnTypedRune(func(r rune) { // This is the event handler for typing characters
		e.HandleRuneInput(r)
        s.textGrid.SetText(e.Buffer.Read())
        s.UpdateCaret()
		s.container.Refresh()
	})

	w.Resize(fyne.NewSize(800, 600))
    w.Show()
    app.Run()
}

// UpdateCaret updates the position of the caret on the screen
func (s *Screen) UpdateCaret() {
    x, y := s.editor.Cursor.GetPosition()
    s.textGrid.SetStyle(x, y, CaretStyle{})
    s.textGrid.Refresh()
}
