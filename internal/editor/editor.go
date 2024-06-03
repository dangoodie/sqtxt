package editor

import (
    "fmt"
    "io/ioutil"
    "os"
)

type Editor struct {
    // define the editor struct here
}

func Start() {
    fmt.Println("Starting editor...")

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

    fmt.Println("File contents:", string(data))
}


