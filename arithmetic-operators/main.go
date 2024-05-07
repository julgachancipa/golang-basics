// https://go.dev/doc/tutorial/getting-started

package main

import "fmt"

func main() {
	// Addition

	x := 50
	y := 10

	result := x + y
	fmt.Println("Addition: ", result)

	// Subtraction
	result = x - y
	fmt.Println("Subtraction: ", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication: ", result)

	// Division
	result = x / y
	fmt.Println("Division: ", result)

	// Module
	result = x % y
	fmt.Println("Module: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

}
