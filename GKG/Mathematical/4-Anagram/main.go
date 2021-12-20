package main

import (
	"fmt"
	"strings"
)

func main() {
	stringA := "geeks"
	stringB := "sekgs"
	if isAnagram(stringA, stringB) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isAnagram(a string, b string) bool {
	isAnagram := true
	if len(a) != len(b) {
		isAnagram = false
	}
	arrA := strings.Split(a, "")
	// fmt.Print(arrA)
	arrB := strings.Split(b, "")
	// fmt.Print(arrB)

	for i := range arrA {
		for j := range arrB {

			if arrA[i] == arrB[j] {

				copy(arrB[j:], arrB[j+1:]) // Shift arrB[i+1:] left one index.
				arrB[len(arrB)-1] = ""     // Erase last element (write zero value).
				arrB = arrB[:len(arrB)-1]  // Truncate slice.
				break
			}
		}
	}
	if len(arrB) > 0 {
		isAnagram = false
	}

	// fmt.Print(arrA == arrB)
	return isAnagram
}
