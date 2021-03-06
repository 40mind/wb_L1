package main

import "fmt"

func main() {
	mass := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapa := make(map[int][]float32)
	for _, elem := range mass {
		mapa[(int(elem)/10)*10] = append(mapa[(int(elem)/10)*10], elem)
	}
	fmt.Println(mapa)
}
