package editor

import (
    "fmt"
)

// HandleInput processes user input
type Position struct {
    row, col int
    
}

func (e *Editor)HandleInput(key rune) {
    switch key {
        case 27: // Escape key

        case 37: // Left arrow
            e.Cursor.row--
        case 38: // Up arrow
            e.Cursor.col--
        case 39: // Right arrow
            e.Cursor.row++
        case 40: // Down arrow  
            e.Cursor.col++

        default:
            fmt.Println("Key pressed:", key)
            //editor.buffer.Insert(editor.Cursor, []byte(string(key))
            //editor.Cursor.x++
    }
}

