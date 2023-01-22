package main

import "fmt"

func main() {
	fmt.Println(occur([]string{"cat", "dog", "elephant"}, "dog"))
}

//Write a function that takes a slice of strings and a string,
//and returns the index of the first occurrence of the string in the slice.
// Input: (["cat", "dog", "elephant"], "dog") Output: 1

func occur(a []string, b string) int {
	c := -1
	for i, v := range a {

		if v == string(b) {
			c = i
			break
		}

	}
	return c
}
