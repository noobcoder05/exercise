package main

import "fmt"

func main() {
	result := largest([]int{1, 2, 3, 4, 5})
	fmt.Println(result)

}

//Write a function that takes a slice of integers and returns the largest number in the slice.
// Input: [1, 2, 3, 4, 5] Output: 5

func largest(l []int) int {

	n := l[0]

	for _, v := range l {

		if v > n {

			n = v
		}

	}
	return n
}
