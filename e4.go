package main

import "fmt"

func main() {

	result := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
}

//Write a function that takes a slice of integers and returns the sum of all the elements in the slice.
// Input: [1, 2, 3, 4, 5] Output: 15
func sum(s []int) int {
	newslice := 0
	for _, v := range s {
		newslice = newslice + v

	}
	return newslice
}
