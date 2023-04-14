package main

import (
	"fmt"
)

func main() {
	// using var
	// var name string = "Kai"
	var age int = 25
	const isCool = true

	// short hand declaration
	// name := "Kai"
	// size := 1.3

	name, size := "Kai", 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)

}
