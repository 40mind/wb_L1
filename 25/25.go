package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println("Start")
	Sleep(5 * time.Second)
	fmt.Println("End")
}
