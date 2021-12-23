package main

import "fmt"

func main() {
	if isAnagramPalindrom("abcbcba") {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

func isAnagramPalindrom(s string) bool {
	var freq [26]int
	for _, chr := range s {
		freq[chr-'a']++
	}
	var countSingle int
	// fmt.Println(freq)
	for _, i := range freq {
		// fmt.Println(i, " - ", countSingle)
		if i%2 != 0 {
			countSingle++
		}
		if countSingle == 2 {
			return false
		}
	}

	return true
}
