package editor

import (
	"bufio"
	"fmt"
	"os"

	//util "github.com/dangoodie/sqtxt/pkg/util"
	buffer "github.com/dangoodie/sqtxt/internal/buffer"
	display "github.com/dangoodie/sqtxt/internal/display"
)

type Editor struct {
    display display.Display
    buffer buffer.Buffer
    cursor Position
    filename string
}

func Start() {
    fmt.Println("Starting editor...")

    if len(os.Args) < 2 {
        fmt.Println("No file specified")
        os.Exit(1)
    }

    filename := os.Args[1]
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }
    

    // load the file contents into the editor
    editor := Editor{
        buffer: buffer.NewBuffer(data),
        cursor: Position{0, 0},
        filename: filename,
        display: display.NewDisplay(),
    }
    fmt.Println("Loaded file:", filename)

    editor.Run()
}

// Run is the main loop of the editor
func (e *Editor) Run() {
    fmt.Println("Running editor...")
    display.Init()
    input := bufio.NewReader(os.Stdin)
    defer Close()

    for {
        e.display.Render(e.buffer)
        key, _, err := input.ReadRune()
        if err != nil {
            fmt.Println("Error reading input:", err)
            break
        }
        e.HandleInput(key, e)
    }
}

func Close() {
    fmt.Println("Closing editor...")
    display.Close()
    os.Exit(0)
}
