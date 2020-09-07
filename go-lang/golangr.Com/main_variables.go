/*
Exercises
Calculate the year given the date of birth and age
Create a program that calculates the average weight of 5 people.
*/

package main

import "fmt"

func main() {
	/*
		age := 26
		dateOfBirth := 1994
		year := age + dateOfBirth
		fmt.Println(year)
	*/

	// AVERAGE WEIGHT OF THE PEOPLE
	w1 := 50
	w2 := 30
	w3 := 20
	w4 := 25
	w5 := 35

	avg := (w1 + w2 + w3 + w4 + w5) / 5
	fmt.Println(avg)

}
