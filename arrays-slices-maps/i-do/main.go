package main

import "fmt"

func main() {
	// ARRAYS
	// create an array arr that holds 7 integers
	var arr[7]int

	// at index four of arr, set the value to 25
	arr[4] = 25

	// declare and initialize an array called names which contains "Peter", "Shelby", and "Zach" in that order
	names := []string{"Peter", "Shelby", "Zach"}

	// SLICES

	// slices are typed only by the elements they contain, not the number of elements
	// declare and initialize a slice of strings named gophers with a length of 5.
	gophers := make([]string, 5)

	// at index 0 of gophers set the value equal to "Mac" and 1 to "Tosh"
	gophers[0] = "Mac"
	gophers[1] = "Tosh"

	// set a slice called rails equal to a slice of names starting at index 1 and ending at the length of names
	//rails := names[1:]

	// append the string "Bobby" into the slice rails
	//append(rails, "Bobby")

	// MAPS
	// (map[key-type]val-type)
	// create a map called cartog which takes strings for keys and float64 for values.
	cartog := make(map[string]float64)

	// within cartog declare the "long" key to equal -86.238333 and the "lat" key to equal 39.203056
	cartog["long"] = -86.238333
	cartog["lat"] = 39.203056

	// use fmt.Println to print the length of names
	fmt.Println(len(names))

	// print out the map cartog
	fmt.Println(cartog)
}
