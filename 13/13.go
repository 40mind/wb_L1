package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Printf("a:%d b:%d\n", a, b)

	// 1 способ
	a, b = b, a
	fmt.Printf("a:%d b:%d\n", a, b)

	// 2 способ
	a += b
	b = a - b
	a -= b
	fmt.Printf("a:%d b:%d\n", a, b)

	// 3 способ
	a *= b
	b = a / b
	a /= b
	fmt.Printf("a:%d b:%d\n", a, b)
}
