package main

import "fmt"

func main() {
	// Strict value
	var name string

	name = "Testing Variable"
	fmt.Println(name)

	// Loose value
	var value = "Nama Variabel"
	fmt.Println(value)

	// Define data type
	var age int8 = 30
	//  or not
	var ages = 30
	fmt.Println(age, ages)

	// Without var
	country := "Indonesia"
	fmt.Println(country)
}
