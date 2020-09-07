/*
Exercises
Create an slices with the number 0 to 10
Create an slices of strings with names
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reflect.TypeOf(a))
}
