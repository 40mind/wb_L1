package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Printf("%d ", <-ch)
		}
	}()

	time.Sleep(5 * time.Second)
}
