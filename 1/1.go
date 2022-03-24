package main

import "fmt"

type Human struct {
	age int
}

func (*Human) foo() {
	fmt.Println("bar")
}

type Action struct {
	Human
}

func main() {
	ac := &Action{}
	ac.foo()
}
