package main

import "fmt"


func main () {
    var(
            a = func() (string) {
                return "I am a function without name!"
            }
       )
    var str string
    str = a()
    fmt.Println(str)
}
