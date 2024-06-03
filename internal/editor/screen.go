package editor

import (
    "fmt"
    "golang.org/x/term"
    "os"
)


type Size struct {
    width int
    height int
}

// InitializeScreen sets up the terminal screen
func Init() {
    // Code to initialize the screen
    // Check if the terminal supports ANSI escape codes
    if !term.IsTerminal(int(os.Stdout.Fd())) {
        fmt.Println("Terminal does not support ANSI escape codes")
        os.Exit(1)
    }


    // Hide the cursor
    fmt.Print("\033[?25l")
    // Clear the screen
    fmt.Print("\033[2J")
    // Move the cursor to the top-left corner
    fmt.Print("\033[H")

    // Set the terminal to raw mode
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println("Error setting terminal to raw mode:", err)
        os.Exit(1)
    }

    _ = oldState
}

func Close() {
    // Show the cursor
    fmt.Print("\033[?25h")
}

func GetSize() Size {
    // Get the size of the terminal
    width, height, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        fmt.Println("Error getting terminal size:", err)
        os.Exit(1)
    }

    return Size{width, height}
}
