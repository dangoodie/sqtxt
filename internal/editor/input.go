package editor

import (
	"fmt"
	structs "github.com/dangoodie/sqtxt/internal/structs"
)

type Cursor struct {
	curCol   int
	Position structs.Position
}

// NewCursor creates a new cursor
func NewCursor() Cursor {
	return Cursor{0, structs.Position{0, 0}}
}

func (c *Cursor) GetPosition() (int, int) {
	return c.Position.Row, c.Position.Col
}

func (e *Editor) HandleKeyInput(key string) {
	switch key {
	case "Escape":
		e.Quit()
	case "Left": // Left arrow
		e.CursorLeft()
	case "Up": // Up arrow
		e.CursorUp()
	case "Right": // Right arrow
		e.CursorRight()
	case "Down": // Down arrow
		e.CursorDown()
	case "BackSpace":
		x, y := e.Cursor.GetPosition()
		delPos := structs.Position{x, y - 1}
		e.Buffer.Delete(delPos)
		e.CursorLeft()
	default:
		fmt.Println("Unknown key:", key)
	}
}

func (e *Editor) HandleRuneInput(r rune) {
	e.Buffer.Insert(e.Cursor.Position, []byte(string(r)))
	e.CursorRight()
	fmt.Println("Rune pressed:", string(r))
}

// CursorUp moves the cursor up
func (e *Editor) CursorUp() {
	if e.Cursor.Position.Row > 0 {
		e.Cursor.Position.Row--
	}

	e.Cursor.Position.Col = e.Cursor.curCol

	if e.Cursor.Position.Col > e.Buffer.LineLength(e.Cursor.Position.Row) {
		e.Cursor.Position.Col = e.Buffer.LineLength(e.Cursor.Position.Row)
	}

	if e.Cursor.Position.Col < 0 {
		e.Cursor.Position.Col = 0
	}
}

// CursorDown moves the cursor down
func (e *Editor) CursorDown() {
	if e.Cursor.Position.Row < e.Buffer.NumLines()-1 {
		e.Cursor.Position.Row++
	}

	e.Cursor.Position.Col = e.Cursor.curCol

	if e.Cursor.Position.Col > e.Buffer.LineLength(e.Cursor.Position.Row) {
		e.Cursor.Position.Col = e.Buffer.LineLength(e.Cursor.Position.Row)
	}

	if e.Cursor.Position.Col < 0 {
		e.Cursor.Position.Col = 0
	}
}

// CursorLeft moves the cursor left
func (e *Editor) CursorLeft() {
	if e.Cursor.Position.Col > 0 {
		e.Cursor.Position.Col--
	}

	e.Cursor.curCol = e.Cursor.Position.Col
}

// CursorRight moves the cursor right
func (e *Editor) CursorRight() {
	if e.Cursor.Position.Col < e.Buffer.LineLength(e.Cursor.Position.Row) {
		e.Cursor.Position.Col++
	}

	e.Cursor.curCol = e.Cursor.Position.Col
}
