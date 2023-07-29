package main

import (
    "fmt"
    "os"
    "os/exec"
    
)

func main() {

    // construct `go version` command
    cmd := exec.Command("lscpu")
    
    // configure `Stdout` and `Stderr`
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    // run command
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err )
    }

}
