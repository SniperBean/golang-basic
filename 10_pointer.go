package main

import "fmt"

func zero (number *int) {
    *number = 0
}

func putchar(char *byte) {
    fmt.Printf("%c", *char)
}

func putstr(str string) {
    for i:=0; i<len(str); i++ {
        char := str[i]
        putchar(&char)
    }
}

func main () {
    A := 5
    zero(&A)

    fmt.Println(A)

    str := "I am a wired test"
    putstr(str)
}
