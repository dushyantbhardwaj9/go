package main

import "fmt"

func main() {
	if IsValidParantheses("({})") {
		fmt.Print("Valid Parantheses ")
	} else {
		fmt.Print("Invaid Parantheses")
	}
}

// IsValidParantheses takes string and returns true
// if the string is valid otherwise return true
func IsValidParantheses(s string) bool {
	if len(s) < 2 {
		return false
	}

	opens := map[rune]struct{}{
		'(': {},
		'{': {},
		'[': {},
	}

	complements := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stackOpen []rune

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if _, ok := opens[r]; ok {
			stackOpen = append(stackOpen, r)
			continue
		}

		if len(stackOpen) == 0 {
			return false
		}

		if stackOpen[len(stackOpen)-1] != complements[r] {
			return false
		}

		stackOpen = stackOpen[:len(stackOpen)-1]
	}

	return len(stackOpen) == 0
}
