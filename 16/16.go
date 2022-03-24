package main

import "fmt"

func QuickSort(mass []int) {
	if len(mass) <= 1 {
		return
	}

	center := Part(mass)

	QuickSort(mass[center:])
	QuickSort(mass[:center])
}

func Part(mass []int) int {
	center := mass[len(mass)/2]

	left := 0
	right := len(mass) - 1

	for {
		for ; mass[left] < center; left++ {

		}

		for ; mass[right] > center; right-- {

		}

		if left >= right {
			return right
		}

		mass[left], mass[right] = mass[right], mass[left]
	}
}

func main() {
	mass := []int{1, 2, 5, 9, 7, 3}

	QuickSort(mass)
	fmt.Println(mass)
}
