package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(slargest([]int{1, 2, 3, 4, 5}))

}

// Write a function that takes a slice of integers and returns the second-largest number in the slice.
// Input: [1, 2, 3, 4, 5] Output: 4
func slargest(s []int) int {

	sort.Ints(s)
	return s[len(s)-2]
}
