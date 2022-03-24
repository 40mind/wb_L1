package main

import "fmt"

func main() {
	var set1 = map[int]bool{5: true, 10: true, 3: true, 0: true}
	var set2 = map[int]bool{7: true, 0: true, 5: true, 1: true}
	set3 := make(map[int]bool)

	for elem := range set1 {
		if _, ok := set2[elem]; ok {
			set3[elem] = true
		}
	}

	for elem := range set3 {
		fmt.Printf("%d ", elem)
	}
}
