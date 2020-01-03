package main

import (
        "fmt"
        "time"
       )

func main () {
    // ref => https://golang.org/src/time/format.go
    // oh... Come on! seriously? what's wrong with you Golang?  
    fmt.Println(time.Now().Format("2006-1-2 03:04:00"))
    fmt.Println(time.Now().Format("Mon, Jan 2, 2006 at 3:04pm"))
}
