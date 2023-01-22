package main

import "fmt"

func main() {

	result := asc([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
}

//Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order.
//Input: [1, 2, 3, 4, 5] Output: true
func asc(a []int) bool {

	for i := 0; i < len(a)-1; i++ {

		if a[i] < a[i+1] {

			continue
		} else {
			return false
		}

	}
	return true
}
