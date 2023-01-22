package main

import "fmt"

func main() {

	result := big([]string{"cat", "dog", "elephant", "lion"})
	fmt.Println(result)

}

//Write a function that takes a slice of strings and returns a new slice with all strings that have more than 5 characters.
// Input: ["cat", "dog", "elephant", "lion"] Output: ["elephant"]

func big(b []string) []string {
	n := []string{}
	for _, v := range b {

		if len(v) > 5 {
			n = append(n, v)
		}
	}
	return n
}
