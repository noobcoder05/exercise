package main

import "fmt"

func main() {

	result := even([]int{1, 2, 3, 4, 5})
	fmt.Println(result)

}

//Write a function that takes a slice of integers and returns a new slice with only the even numbers.
//Input: [1, 2, 3, 4, 5] Output: [2, 4]
func even(e []int) []int {

	slice := []int{}
	for _, v := range e {

		if v%2 == 0 {

			slice = append(slice, v)

		}

	}
	return slice
}
