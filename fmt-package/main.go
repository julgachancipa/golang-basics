// https://go.dev/doc/tutorial/getting-started

package main

import "fmt"

func main() {

	// Println
	helloMessage := "Hello"
	worldMessage := "world"
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Emma"
	age := 23
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%v is %v years old\n", name, age)

	// Sprintf
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message)

	// Data types
	fmt.Printf("name: %T\n", name)
	fmt.Printf("name: %T\n", age)
}
