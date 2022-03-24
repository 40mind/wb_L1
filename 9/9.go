package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	mass := [...]int{5, 12, 0, -1, 109, -34, 2}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, item := range mass {
			ch1 <- item
		}
	}()

	go func() {
		for item := range ch1 {
			ch2 <- item * 2
		}
	}()

	go func() {
		for item := range ch2 {
			fmt.Println(item)
		}
	}()

	ch3 := make(chan os.Signal, 1)
	signal.Notify(ch3, os.Interrupt)
	<-ch3
}
