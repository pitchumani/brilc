package main

import (
    "fmt"
    "os"
)

func compile_file (fname string) {
    fmt.Println ("Compiling file: ", fname)
}

func main () {
    var args = []string(os.Args[0:])
    argc := len(args)
    if argc == 1 {
        fmt.Println ("Usage: ", args[0], " <bril-json-file> [ bril-json-file ]*")
        os.Exit(1)
    }

    for i := 1; i < len(args); i++ {
        compile_file (args[i])
    }
}
