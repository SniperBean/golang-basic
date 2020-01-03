package main

import "fmt"

func main () {
    // Array
    var array[2] string

    array[0] = "hello"
    array[1] = "world"

    fmt.Println("Print Array")
    fmt.Println(array[0] + " " + array[1] + "!")

    fmt.Println()

    // Slice
    slice1 := []string {"hello", "world"}
    fmt.Println("Print Slice with index")
    fmt.Println(slice1[0] + " " + slice1[1] + "!")

    slice2 := []int {1, 2, 3, 4, 5, 6}
    fmt.Println("Print Slice with range")
    fmt.Println(slice2[0:4])
    fmt.Println(slice2[2:5])
    fmt.Println(slice2[:5])
    fmt.Println(slice2[3:])

    fmt.Println()

    // MAP
    data := make(map[string]string)

    data["username"] = "Bean"
    data["password"] = "Lean"
    fmt.Println("Print Map")
    fmt.Println("{ Username: " + data["username"] + " Password: " + data["password"] + " }")

    // Interface
    mixedData := make(map[string]interface{})
    mixedData["username"] = "Dean"
    mixedData["password"] = 12345678
    fmt.Println("Print Map with Interface")
    fmt.Println("{ Username:", mixedData["username"], "Password", mixedData["password"], "}")

    var mixed interface{}
    fmt.Println("Print Interface")
    mixed = 123
    fmt.Println(mixed)
    mixed = "Disco Dingo"
    fmt.Println(mixed)
    mixed = []string{"A", "B", "C"}
    fmt.Println(mixed)
}
