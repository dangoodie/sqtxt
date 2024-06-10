package editor

import (
	"fmt"
	"os"

	//util "github.com/dangoodie/sqtxt/pkg/util"
	buffer "github.com/dangoodie/sqtxt/internal/buffer"
    structs "github.com/dangoodie/sqtxt/internal/structs"
)

type Editor struct {
	Buffer   *buffer.Buffer
	Cursor   structs.Position
	Filename string
}

func Start() Editor {
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
		Buffer:   buffer.NewBuffer(data),
		Cursor:   structs.Position{0, 0},
		Filename: filename,
	}
	fmt.Println("Loaded file:", filename)

	return editor
}
