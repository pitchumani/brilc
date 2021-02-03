package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)

func compile_function (function map[string]interface{}) {
    fmt.Println ("Function: ", function["name"].(string))
    // ignore args for now
    // process instructions (array of json objects)
    instrs := function["instrs"].([]interface{})
    insnCount := len(instrs)
    fmt.Println ("Total insns: ", insnCount)
    for insnIndex := 0; insnIndex < insnCount; insnIndex++ {
        fmt.Println ("insn ", insnIndex, ":")
        // each instruction is a set of json objects (can map with key, value)
        insn := instrs[insnIndex].(map[string]interface{})
        for ikey, ival := range insn {
            fmt.Println ("  attribs ", ikey, ":", ival)
        }
    }
}

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
    for i := 0; i < len(functions); i++ {
        compile_function (functions[i].(map[string]interface{}))
    }
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
