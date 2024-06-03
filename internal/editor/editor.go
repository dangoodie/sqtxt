package editor

import (
    "fmt"
    "io/ioutil"
    "os"

    //display "github.com/dangoodie/internal/display"
    //util "github.com/dangoodie/pkg/util"

)

type Editor struct {
    buffer Buffer
    cursor int
    filename string
    //display display
    //util util
}

func Start() {
    //fmt.Println("Starting editor...")

    if len(os.Args) < 2 {
        fmt.Println("No file specified")
        os.Exit(1)
    }

    filename := os.Args[1]
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }
    // load the file contents into the editor
    editor := Editor{
        buffer: Buffer{data: data},
        cursor: 0,
        filename: filename,
    }
}

