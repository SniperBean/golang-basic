package main

import "fmt"

func main () {
    data := map[string]string{
              "username": "follow",
              "password": "moon"}

    _, exists := data["money"]

    if exists {
        fmt.Println("Still not exists cause moon lies")
    } else {
        fmt.Println("It doesn't exist")
    }
}
