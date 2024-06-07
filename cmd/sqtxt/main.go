package main

import (
    editor "github.com/dangoodie/sqtxt/internal/editor"
    display "github.com/dangoodie/sqtxt/internal/display"
)

func main() {
    e := editor.Start()
    display.Start(e)
}
