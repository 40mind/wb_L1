package main

import "fmt"

func BinSearch(mass []int, a int) int {
	left := 0
	right := len(mass) - 1
	for left < right {
		center := (right + left) / 2
		if mass[center] >= a {
			right = center
		} else {
			left = center + 1
		}
	}
	return left
}

func main() {
	mass := []int{1, 2, 4, 6, 7, 8, 11}
	fmt.Printf("%d - индекс нужного элемента", BinSearch(mass, 2))
}
