package editor

import (
    "fmt"
    "os"
    "bufio"

    //util "github.com/dangoodie/sqtxt/pkg/util"

)

type Editor struct {
    buffer Buffer
    cursor Position
    filename string
    size Size
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
    
    // Get the size of the terminal
    size := GetSize()

    // load the file contents into the editor
    editor := Editor{
        buffer: NewBuffer(data),
        cursor: Position{0, 0},
        filename: filename,
        size: size,
    }
    fmt.Println("Loaded file:", filename)

    editor.Run()
}

// Run is the main loop of the editor
func (editor *Editor) Run() {
    fmt.Println("Running editor...")
    Init()
    input := bufio.NewReader(os.Stdin)
    defer Close()

    for {
        Render(editor)
        key, _, err := input.ReadRune()
        if err != nil {
            fmt.Println("Error reading input:", err)
            break
        }
        editor.HandleInput(key, editor)
    }
}
