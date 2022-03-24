package main

import "fmt"

func main() {
	str := "главрыба"
	runemass := []rune(str)

	for i := 0; i < len(runemass)/2; i++ {
		runemass[i], runemass[len(runemass)-i-1] = runemass[len(runemass)-i-1], runemass[i]
	}

	str = string(runemass)
	fmt.Println(str)
}
