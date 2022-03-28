package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mx    sync.Mutex
}

func Increment(cnt *Counter, wg *sync.WaitGroup) {
	cnt.mx.Lock()
	cnt.count++
	cnt.mx.Unlock()
	wg.Done()
}

func main() {
	cnt := &Counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Increment(cnt, &wg)
	}

	wg.Wait()
	fmt.Println(cnt.count)
}
