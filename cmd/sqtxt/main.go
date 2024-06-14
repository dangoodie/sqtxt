package main

import (
	display "github.com/dangoodie/sqtxt/internal/display"
	editor "github.com/dangoodie/sqtxt/internal/editor"
)

func main() {
	e := editor.Start()
	display.Start(e)
}
