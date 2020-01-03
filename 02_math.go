package main

import (
	"fmt"
	"math"
	"math/rand"
       )

func swap(x, y string) (string, string) {
	return y, x
}

func main () {
	fmt.Println("This is a random number ", rand.Intn(10000))
	fmt.Println("This is a number of square root of 81", math.Sqrt(81))
	fmt.Println("This is a number of pi", math.Pi)
	a, b := swap("swap", "string")
	fmt.Println("swap swap string to be ", a, b)
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java)
}
