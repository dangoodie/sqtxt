package display

import (
	"fmt"
	"strings"
    buffer "github.com/dangoodie/sqtxt/internal/buffer"
)

// Display represents the terminal display
type Display struct {
    size Size
}
// New Display function
func NewDisplay() Display {
    return Display{size: GetSize()}
}

// Render renders the current state of the editor to the screen
func (d *Display) Render(buffer buffer.Buffer) {
	fmt.Print("\033[2J") // Clear the screen
	lines := strings.Split(string(buffer.Read()), "\n")
    
    for _, line := range lines {
        fmt.Println(line)
    }
}
