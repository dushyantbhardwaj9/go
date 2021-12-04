package main

import "fmt"

func main() {
	rstring := "gnirtSelpmaS"
	fmt.Print(reverse(rstring))
}

func reverse(rstring string) string {
	rns := []rune(rstring) // Convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
