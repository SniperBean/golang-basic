package main

import (
        "fmt"
        "errors"
       )

func numberErrorHandler(paid int) (int, error) {
    if paid < 0 {
        return 0, errors.New("paid: It is not positive")
    }
    return paid, nil
}

func main () {
    if paid, err := numberErrorHandler(-123); err != nil {
        fmt.Println(paid, err)
    }
}
