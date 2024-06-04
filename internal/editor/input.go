package editor

// HandleInput processes user input
type Position struct {
    x int
    y int
}

func (e *Editor)HandleInput(key rune, editor *Editor) {
    switch key {
        case 27: // Escape key
            Close()
        case 37: // Left arrow
            editor.cursor.x--
        case 38: // Up arrow
            editor.cursor.y--
        case 39: // Right arrow
            editor.cursor.x++
        case 40: // Down arrow  
            editor.cursor.y++

        default:
            //editor.buffer.Insert(editor.cursor, []byte(string(key))
            //editor.cursor.x++
    }
}

