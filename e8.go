package main

import "fmt"

func main() {
	result := Removeduplicate([]int{1, 2, 3, 2, 4, 5, 5})
	fmt.Println(result)

}

//Write a function that takes a slice of integers and returns a new slice with all duplicates removed.
// Input: [1, 2, 3, 2, 4, 5, 5] Output: [1, 3, 4]

func Removeduplicate(slice []int) []int {

	newMap := map[int]bool{}
	newslice := []int{}
	for _, value := range slice {

		if _, ok := newMap[value]; !ok {

			newMap[value] = true
			newslice = append(newslice, value)
		}

	}
	return newslice
}
