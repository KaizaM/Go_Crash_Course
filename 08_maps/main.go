package main

import "fmt"

func main() {
	// map is a key value pair

	// Define Map
	emails := make(map[string]string)

	// Assign Key
	emails["Bob"] = "bob@gmail.com"
	emails["Andy"] = "andy@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(emails["Andy"])

}
