package main

import "fmt"

func main() {
	n := 5
	printPattern(n)
}

func printPattern(n int) {
	for i := 1; i < (n + 1); i++ {
		for j := n; j > 0; j-- {
			for k := 0; k < i; k++ {
				fmt.Printf("%d", j)
			}
		}
		fmt.Println()
	}
}
