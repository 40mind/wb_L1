package main

import (
	"github.com/40mind/wb_L1/24/pkg"
)

func main() {
	point1 := pkg.New(10, 5)
	point2 := pkg.New(20, 30)
	pkg.Dist(point1, point2)
}
