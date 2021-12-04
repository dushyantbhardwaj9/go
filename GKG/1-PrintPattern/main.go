package main

import "fmt"

func main() {
	n := 5
	for o := 1; o < (n + 1); o++ {
		for i := n; i > 0; i-- {
			for p := 0; p < o; p++ {
				fmt.Printf("%d", i)
			}
		}
		fmt.Println()
	}
}
