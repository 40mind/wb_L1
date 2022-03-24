package main

import (
	"fmt"
)

func main() {
	// Первый способ: создание канала остановки
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				// …
			}
		}
	}()
	// …
	close(quit)

	// Второй способ: использование одного канала для передачи данных и остановки горутины
	// Для остановки горутины необходимо закрыть канал
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
}

func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}
