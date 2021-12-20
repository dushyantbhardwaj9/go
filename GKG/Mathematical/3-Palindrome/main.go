package main

import "fmt"

func main() {
	rstring := "abcba"
	if isPalindrome(rstring) {
		fmt.Println(rstring, " is a Palindrome")
	} else {
		fmt.Println(rstring, " is NOT a Palindrome")
	}
}

func isPalindrome(rstring string) bool {
	//
	rns := []rune(rstring) // Convert to rune
	isPalindrome := true
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// compare the first and last letter
		if rns[i] == rns[j] {
			continue
		} else {
			isPalindrome = false
		}
		// rns[i], rns[j] = rns[j], rns[i]
	}
	return isPalindrome
}
