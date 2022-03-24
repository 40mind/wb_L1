package main

import (
	"fmt"
	"sync"
)

func main() {
	mapa := make(map[int]int)
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mx.Lock()
			mapa[i] = i * i
			mx.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(mapa[9])
}
