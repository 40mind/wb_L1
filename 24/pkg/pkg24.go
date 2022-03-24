package pkg

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func New(val1, val2 int) *Point {
	return &Point{
		x: val1,
		y: val2,
	}
}

func Dist(p1, p2 *Point) {
	fmt.Println(math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2)))
}
