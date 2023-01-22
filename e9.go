package main

import "fmt"

func main() {

	result := average([]int{1, 2, 3, 4, 5})
	fmt.Println(result)

}

//Write a function that takes a slice of integers and returns the average of all the elements in the slice.
// Input: [1, 2, 3, 4, 5] Output: 3

func average(a []int) int {

	newslice := 0
	av := 0

	for _, v := range a {

		newslice = newslice + v

	}
	av = (newslice) / len(a)
	return av
}
