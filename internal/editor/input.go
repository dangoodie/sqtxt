package editor

import (
    "fmt"
)

func (e *Editor)HandleInput(key rune) {
    switch key {
        case 37: // Left arrow
            e.Cursor.Row--
        case 38: // Up arrow
            e.Cursor.Col--
        case 39: // Right arrow
            e.Cursor.Row++
        case 40: // Down arrow  
            e.Cursor.Col++

        default:
            fmt.Println("Key pressed:", key)
            e.Buffer.Insert(e.Cursor, []byte{byte(key)})
            e.Cursor.Col++
    }
}

