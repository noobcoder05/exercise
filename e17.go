package main

import "fmt"

func main() {

	fmt.Println(given([]string{"cat", "dog", "elephant"}, "e"))
}

//Write a function that takes a slice of strings and returns a new slice with all strings that start with a given letter.
//Input: (["cat", "dog", "elephant"], "e") Output: ["elephant"]

func given(b []string, a string) []string {

	c := []string{}

	for _, v := range b {

		if v[0] == a[0] {

			c = append(c, v)
		}

	}
	return c

}
