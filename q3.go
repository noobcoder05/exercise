package main

import "fmt"

func main() {
	lovely := multi(map[string]int{"a": 1, "b": 2, "c": 3})

	fmt.Println(lovely)

}

//Write a function that takes a map and returns a new map with all the values multiplied by 2.
//For example, given the input map[string]int{"a": 1, "b": 2, "c": 3},
//the function should return a new map map[string]int{"a": 2, "b": 4, "c": 6}.
func multi(lovely map[string]int) map[string]int {
	for k, v := range lovely {
		lovely[k] = 2 * v

	}
	return lovely
}
