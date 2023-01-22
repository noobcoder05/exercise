package main

import "fmt"

func main() {

	fmt.Println(occ([]int{1, 2, 3, 4, 5}, 2))

}

//Write a function that takes a slice of integers and returns the index of the first occurrence of a given number.
// Input: ([1, 2, 3, 2, 4, 5], 2) Output: 1

func occ(o []int, c int) int {
	var count int
	for _, v := range o {

		if v == c {
			count++
		}

	}
	return count
}
