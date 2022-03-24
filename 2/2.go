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
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println()

	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(mass[i] * mass[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func multi(mass []int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- mass[i] * mass[i]
	}
}
