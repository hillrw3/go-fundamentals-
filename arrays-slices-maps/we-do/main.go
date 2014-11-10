package main

import "fmt"

func main() {
	// create an array called climateCo filled with the following strings: "Zone 4", "Zone 5", "Zone 6",
	// "Zone 7", and "Zone 8", but make the length of the array 6
	climateCo := [6]string{"Zone 4", "Zone 5", "Zone 6", "Zone 7", "Zone 8"}

	// print the length of the array followed by Denver's climate zone (Zone 5)
	fmt.Println(len(climateCo), climateCo[1])

	// create a slice of climateCo called plainsClimate which is a slice including climateCo[0] and climateCo[1]
	plainsClimate := climateCo[:2]

	// print plainsClimate
	fmt.Println(plainsClimate)

	// create a slice of integers called sliceInt with a length of 2
	sliceInt := make([]int, 2)

	// print the length of sliceInt
	fmt.Println(len(sliceInt))

	// in one line, create a map named dates with a key "year_made" with a value of 2007, and "year_updated"
	// with the value of 2014
	dates := map[string]int{"year_made":2007, "year_updated":2014}

	// append into dates the key "major_update" with the value 2011
	dates["major_update"] = 2011

	// set a the name present equal to a boolean that determines whether dates["major_update"] exists
	// use a skip do this in one line
	// _, present := dates["major_update"]

	// print "present:" followed by the boolean you just named
	// fmt.Println("present:", present)

	// delete the key/value pair "major_update" that you recently created
	delete(dates, "major_update")

	// reset present so that it determines if "major_update" exists in dates, then print it again
	_, present := dates["major_update"]
	fmt.Println("present:", present)

}
