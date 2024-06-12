package editor

import (
	"fmt"
    structs "github.com/dangoodie/sqtxt/internal/structs"
)

func (e *Editor) HandleKeyInput(key string) {
	switch key {
	case "Escape":
		fmt.Println("Quitting...")
		e.Quit()
	case "Left": // Left arrow
		e.Cursor.Col--
	case "Up": // Up arrow
		e.Cursor.Row--
	case "Right": // Right arrow
		e.Cursor.Col++
	case "Down": // Down arrow
        e.Cursor.Row++
    case "BackSpace":
        delPos := structs.Position{e.Cursor.Row, e.Cursor.Col - 1}
        e.Buffer.Delete(delPos)
        e.Cursor.Col--
    default:
        fmt.Println("Unknown key:", key)
	}
}

func (e *Editor) HandleRuneInput(r rune) {
	e.Buffer.Insert(e.Cursor, []byte(string(r)))
	e.Cursor.Col++
	fmt.Println("Rune pressed:", string(r))
}
