package display

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	editor "github.com/dangoodie/sqtxt/internal/editor"
)

// InitializeScreen sets up the terminal screen
func Start(e editor.Editor) {
	// Code to initialize the screen
	fmt.Println("Initializing screen...")
	app := app.New()
	window := app.NewWindow("Hello")

	content := fmt.Sprintf("Filename: %s\nBuffer: %s", e.Filename, string(e.Buffer.Read()))
	text := widget.NewLabel(content)

	window.SetContent(text)
	window.ShowAndRun()
}
