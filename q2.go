package main

import (
	"fmt"
	"sort"
)

func main() {
	res := sorted(map[string]int{"a": 1, "c": 2, "b": 3})
	fmt.Println(res)
}

// Write a function that takes a map and returns a slice of the keys in the map, sorted in alphabetical order.
// For example, given the input map[string]int{"a": 1, "c": 2, "b": 3}, the function should return a slice []string{"a", "b", "c"}.
func sorted(baba map[string]int) []string {
	slice := []string{}
	for key := range baba {
		slice = append(slice, key)

	}
	sort.Strings(slice)
	return slice

}
