package main

import "fmt"

func test() (string, int) {
	return "Bean", 123456
}

func main() {
    username, time := test()
    fmt.Println(username, time)
}
