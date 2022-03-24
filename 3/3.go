package main

import (
	"fmt"
	"sync"
)

func main() {
	mass := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)
	defer close(ch)
	go multi(mass, ch)
	var sum int
	for i := 0; i < 5; i++ {
		sum += <-ch
	}
	fmt.Println(sum)

	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5)
	sum = 0
	for i := 0; i < 5; i++ {
		go func(i int) {
			mx.Lock()
			sum += mass[i] * mass[i]
			mx.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}

func multi(mass []int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- mass[i] * mass[i]
	}
}
