package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func Workers(n int, ch chan int) {
	for i := 0; i < n; i++ {
		go func(ch chan int, i int) {
			for item := range ch {
				fmt.Printf("Worker %d, message %d\n", i, item)
			}
		}(ch, i)
	}
}

func WriteData(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	for {
		ch <- rand.Intn(100)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan int)

	Workers(5, ch1)
	go WriteData(ch1)

	ch2 := make(chan os.Signal, 1)
	signal.Notify(ch2, os.Interrupt)
	<-ch2
}
