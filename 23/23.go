package main

import "fmt"

func DelI(mass []int, i int) []int {
	copy(mass[i:], mass[i+1:])
	mass = mass[:len(mass)-1]
	return mass
}

func main() {
	mass := []int{9, 12, 0, 3, 8, 10, 45, 2}
	mass = DelI(mass, 3)
	fmt.Println(mass)
}
