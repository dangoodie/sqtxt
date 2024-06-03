package util

import "log"

// Logger is a simple logger
var Logger = log.New(os.Stdout, "LOG: ", log.LstdFlags)

