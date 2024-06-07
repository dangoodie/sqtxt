package editor

// HandleInput processes user input
type Position struct {
    row, col int
    
}

func (e *Editor)HandleInput(key rune, editor *Editor) {
    switch key {
        case 27: // Escape key

        case 37: // Left arrow
            editor.Cursor.row--
        case 38: // Up arrow
            editor.Cursor.col--
        case 39: // Right arrow
            editor.Cursor.row++
        case 40: // Down arrow  
            editor.Cursor.col++

        default:
            //editor.buffer.Insert(editor.Cursor, []byte(string(key))
            //editor.Cursor.x++
    }
}

