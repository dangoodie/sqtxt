package display

import (
	"fmt"
	"image/color"

    util "github.com/dangoodie/sqtxt/pkg/util"

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
	cObj = append(cObj, s.caret.Rectangle)
	for _, text := range s.textObj {
		cObj = append(cObj, text)
	}
	s.container = container.NewWithoutLayout(cObj...)

	return s
}

func (s *Screen) loadText() {
    // get configuration
    cfg := util.GetConfig()

	s.textObj = []*canvas.Text{}
	for i, line := range s.editor.Buffer.Read() {
		text := canvas.NewText(string(line), color.White)
		text.TextStyle.Monospace = true
        text.TextSize = float32(cfg.FontSize)

		text.Move(fyne.NewPos(0, float32(i*cfg.RowHeight)))
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
	window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) { // This is the event handler for key presses
		e.HandleKeyInput(string(ev.Name))
		screen.loadText()
		screen.caret.Move(e.Cursor)
		screen.container.Refresh()
	})

	window.Canvas().SetOnTypedRune(func(r rune) { // This is the event handler for typing characters
		e.HandleRuneInput(r)
		screen.loadText()
		screen.caret.Move(e.Cursor)
		screen.container.Refresh()
	})

	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
