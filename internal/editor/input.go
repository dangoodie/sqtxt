package editor

import (
	"fmt"
)

func (e *Editor) HandleKeyInput(key string) {
	switch key {
	case "Escape":
		fmt.Println("Escape key pressed")
		e.Quit()
	case "Left": // Left arrow
		fmt.Println("Left arrow")
		e.Cursor.Col--
	case "Up": // Up arrow
		fmt.Println("Up arrow")
		e.Cursor.Row--
	case "Right": // Right arrow
		fmt.Println("Right arrow")
		e.Cursor.Col++
	case "Down": // Down arrow
		fmt.Println("Down arrow")
        e.Cursor.Row++
	}
}

func (e *Editor) HandleRuneInput(r rune) {
	e.Buffer.Insert(e.Cursor, []byte(string(r)))
	e.Cursor.Col++
	fmt.Println("Rune pressed:", string(r))
}
