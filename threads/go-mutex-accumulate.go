//
// go-mutex-accumulate.go
// a critical race between go-routines
// last-update:
//		25 sep 2026 -bjr; created
//
//


package main

import (
	"fmt"
	"time"
)

var accumulate int = 0

func accumulator(lock chan int, completion chan int) {
	<-lock
	i := accumulate 
	fmt.Println("thread sleeping on lock")
	time.Sleep(time.Second)
	accumulate = i+1
	lock <- 1
	completion <- 1 
}

func main() {
	completion := make(chan int)
	lock := make(chan int, 1)
	// channel capacity of 1 lets me go on after this statement
	lock <- 1
	
	for i := 0; i < 5; i++ {
		go accumulator(lock, completion)
	}
	
	for i := 0; i < 5; i++ {
		// the write to the completion channel blocks
		// until there is an accumulator to read the channel
		<-completion
		fmt.Println("an accumulator thread wrote to the completion channel")
	}

	fmt.Println("final accumulate value:", accumulate)
}
