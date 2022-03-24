package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun road web"
	massStr := strings.Split(str, " ")

	for i := 0; i < len(massStr)/2; i++ {
		massStr[i], massStr[len(massStr)-i-1] = massStr[len(massStr)-i-1], massStr[i]
	}

	str = strings.Join(massStr, " ")
	fmt.Println(str)
}
