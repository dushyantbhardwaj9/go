package main

import "fmt"

func main() {
	fmt.Print("Running romanToInt...")
}

// test romanToInt
func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := romanMap[rune(s[0])]
	for i := 1; i < len(s); i++ {

		roman := rune(s[i])
		prev := romanMap[rune(s[i-1])]
		if prev < romanMap[roman] {

			sum -= (prev * 2)
		}

		sum += romanMap[roman]
		// fmt.Printf("%c -%d, %d, %d ---", roman, romanMap[roman], prev, sum)
	}
	return sum
}
