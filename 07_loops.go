package main

import "fmt"

func main() {
    // Goto
    fmt.Println("use GOTO skip someting, but remember... never use it")
    goto portal
    fmt.Println("Skip me!")

    portal:
    fmt.Println("It's my turn.")

    // For
    fmt.Println("use FOR loop to print numbers from 0 to 3")
    for i := 0; i <= 3; i++ {
        fmt.Println(i)
    }

    // Foreach
    data := []string {"A", "B", "C"}
    fmt.Println("use FOREACH loop to print alphabet from A to C with index and value")
    for index, value := range data {
        fmt.Printf("%d: %s\n", index, value);
    }
    fmt.Println("use FOREACH loop to print alphabet from A to C with index")
    for index := range data {
        fmt.Printf("%d\n", index);
    }
    fmt.Println("use FOREACH loop to print alphabet from A to C with value")
    for _, value := range data {
        fmt.Printf("%s\n", value);
    }

    // While
    i := 0

    fmt.Println("use WHILE loop to print alphabet from 0 to 2")
    for i < 3 {
        i++
        fmt.Println(i)
    }
}
