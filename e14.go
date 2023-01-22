package main

import "fmt"

func main() {

	result := reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
}

//Write a function that takes a slice of integers and returns a new slice with the elements in reverse order.
//Input: [1, 2, 3, 4, 5] Output: [5, 4, 3, 2, 1]
func reverse(b []int) []int {

	a := []int{}
	for i := len(b); i > 0; i-- {
		a = append(a, b[i-1])

	}

	return a
}
