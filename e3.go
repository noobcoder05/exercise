package main

import "fmt"

func main() {

	result := freq("hello")
	fmt.Println(result)

}

//Write a function that takes a string and returns a map with the frequency count of each character in the string.
//Input: "hello" Output: {"h": 1, "e": 1, "l": 2, "o": 1}

func freq(s string) map[string]int {

	m := make(map[string]int)
	for _, v := range s {

		value, ok := m[string(v)]
		if ok {

			m[string(v)] = value + 1
		} else {
			m[string(v)] = 1
		}

	}
	return m
}
