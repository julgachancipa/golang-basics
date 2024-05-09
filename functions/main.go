package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func multipleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello world")
	multipleArguments(1, 2, "Hello there!")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value 1: ", value1)
	fmt.Println("Value 1: ", value2)

	// Ignore value
	value1, _ = doubleReturn(2)
}
