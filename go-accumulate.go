package main

import (
	"fmt"
	"time"
)

var accumulate int = 0

func accumulator(lock, unlock chan int) {
	<-lock
	fmt.Println("thread sleeping on lock")
	time.Sleep(time.Second)
	accumulate += 1
	unlock <- 1
}

func main() {
	lock, unlock := make(chan int), make(chan int)
	for i := 0; i < 5; i++ {
		go accumulator(lock, unlock)
	}
	for i := 0; i < 5; i++ {
		lock <- 1
		<-unlock
	}
	fmt.Println("final accumulate value:", accumulate)
}
