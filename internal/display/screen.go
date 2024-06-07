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
    title := fmt.Sprintf("sqtxt - %s", e.Filename)
	window := app.NewWindow(title)

	content := string(e.Buffer.Read())
	text := widget.NewLabel(content)
    text.TextStyle.Monospace = true

	window.SetContent(text)
	window.ShowAndRun()
}
