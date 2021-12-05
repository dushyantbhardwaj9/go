package main

import "fmt"

func main() {
	num := 3
	printTable(num)

}

func printTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Print(i*num, " ")
	}
}
