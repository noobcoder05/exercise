package main

import "fmt"

func main() {

	result := damn(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "b", "d"})
	fmt.Println(result)

}

//Write a function that takes a map and a slice of strings,
//and returns a new slice with all the strings from the input slice that are keys in the map.
//For example, given the input map[string]int{"a": 1, "b": 2, "c": 3},
//the function should return a slice []string{"a", "b"}.
func damn(bro map[string]int, dude []string) []string {

	newslice := []string{}

	for key := range bro {
		for _, value := range dude {

			if key == value {

				newslice = append(newslice, key)
				continue

			}
		}

	}

	return newslice
}
