package main

import (
        "fmt"
        "os"
       )

func main () {
    handler, err := os.Open("test.txt")
    defer handler.Close()

    if (true) {
        // fake error handler
        fmt.Println("FAKE ERROR -", err)
    }
}

