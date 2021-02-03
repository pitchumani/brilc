package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)

func compile_file (fname string) {
    fmt.Println ("Compiling file: ", fname)
    jfile, err := os.Open (fname)
    if err != nil {
        fmt.Println ("Could not open file ", fname)
        os.Exit(1)
    }
    defer jfile.Close()
    byteValue, _ := ioutil.ReadAll(jfile)
    var program map[string]interface{}
    json.Unmarshal(byteValue, &program)
    functions := program["functions"].([]interface{})
    fmt.Println("nfuncs: ", len(functions))
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
