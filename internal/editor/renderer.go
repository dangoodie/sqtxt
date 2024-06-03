package editor

import (
	"fmt"
	"strings"
)

// Render renders the current state of the editor to the screen
func Render(editor *Editor) {
	fmt.Print("\033[2J") // Clear the screen
	lines := strings.Split(string(editor.buffer.Read()), "\n")
	//numLines := len(lines)

	for i := 0; i < editor.size.height; i++ {
		fmt.Println(lines[i])

	}
}
