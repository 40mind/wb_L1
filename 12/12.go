package main

import "fmt"

func main() {
	mass := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, elem := range mass {
		set[elem] = true
	}

	for elem := range set {
		fmt.Println(elem)
	}

}
