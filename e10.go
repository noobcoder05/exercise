package main

import "fmt"

func main() {

	result := justice([]string{"cat", "dog", "elephant"}, "dog")
	fmt.Println(result)

}

//Write a function that takes a slice of strings and a string, and returns true if the string is contained in the slice.
//Input: (["cat", "dog", "elephant"], "dog") Output: true

func justice(a []string, b string) bool {

	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}
