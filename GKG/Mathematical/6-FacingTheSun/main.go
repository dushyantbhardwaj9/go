package main

import "fmt"

func main() {
	buildings := []int{7, 2, 8, 4, 9, 12}
	fmt.Println(buildingFacingTheSun(buildings))
}

func buildingFacingTheSun(a []int) int {
	count := 0
	max := 0
	for _, val := range a {
		if val > max {
			max = val
			count++
		}
	}
	return count
}
