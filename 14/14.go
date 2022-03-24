package main

import (
	"fmt"
	"reflect"
)

func GetType(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}

func main() {
	xInt := 10
	xString := "Hello"
	xBool := true
	xChan := make(chan int)

	fmt.Println(GetType(xInt))
	fmt.Println(GetType(xString))
	fmt.Println(GetType(xBool))
	fmt.Println(GetType(xChan))
}
