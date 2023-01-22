package main

import "fmt"

func main() {

	result := appear([]int{1, 2, 3, 2, 4, 5}, 2)
	fmt.Println(result)

}

//Write a function that takes a slice of integers and a number and returns the number of times the number appears in the slice.
//Input: ([1, 2, 3, 2, 4, 5], 2) Output: 2
func appear(a []int, b int) int {

	count := 0

	for _, v := range a {

		if v == b {

			count++

		}
	}
	return count
}
