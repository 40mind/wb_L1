package main

import (
	"fmt"
	"strings"
)

func check(word string) bool {
	word = strings.ToLower(word)
	mapa := make(map[string]bool)

	for i := 0; i < len(word); i++ {
		if _, ok := mapa[string(word[i])]; ok {
			return false
		} else {
			mapa[string(word[i])] = true
		}
	}
	return true
}

func main() {
	word := "abCcd"
	fmt.Println(check(word))
}
