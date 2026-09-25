package main

import (
	"fmt"
	"time"
)

var accumulate int = 0

func accumulator(lock chan int, done chan int) {
	<-lock
	fmt.Println("thread sleeping on lock")
	time.Sleep(time.Second)
	accumulate += 1
	lock <- 1
	done <- 1 
}

func main() {
	done := make(chan int)
	lock := make(chan int, 1)
	lock <- 1
	
	for i := 0; i < 5; i++ {
		go accumulator(lock, done)
	}
	
	for i := 0; i < 5; i++ {
		<-done
		fmt.Println("received done ", i)
	}
	fmt.Println("final accumulate value:", accumulate)
}
