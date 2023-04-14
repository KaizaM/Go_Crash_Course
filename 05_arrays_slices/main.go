package main

import "fmt"

func main() {
	// // arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])
	fmt.Println(len(fruitArr))

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[0])
	fmt.Println(fruitSlice[1])
	fmt.Println(fruitSlice[0:3])

}
