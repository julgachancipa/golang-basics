// https://go.dev/doc/tutorial/getting-started

package main

import "fmt"

func main() {
	// Const declaration

	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Int vars declaration

	length := 12
	var height int = 14
	var area int

	fmt.Println(length, height, area)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("int zero value: ", "'", a, "'")
	fmt.Println("float64 zero value: ", "'", b, "'")
	fmt.Println("string zero value: ", "'", c, "'")
	fmt.Println("bool zero value: ", "'", d, "'")

	// Square area

	const squareSide = 10
	squareArea := squareSide * squareSide
	fmt.Println("Square area: ", squareArea)

}
