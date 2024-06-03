package util

import (
    "log"
    "os"
)

// Logger is a simple logger
var Logger = log.New(os.Stdout, "LOG: ", log.LstdFlags)

