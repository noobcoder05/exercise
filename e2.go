package main

import (
	"fmt"
	"sort"
)

func main() {

	result := sortme([]string{"zebra", "monkey", "aardvark"})

	fmt.Println(result)

}

// Write a function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order.
// Input: ["zebra", "monkey", "aardvark"] Output: ["aardvark", "monkey", "zebra"]
func sortme(s []string) []string {

	slice := []string{}

	for _, v := range s {
		slice = append(slice, v)
	}
	sort.Strings(slice)
	return slice

}
