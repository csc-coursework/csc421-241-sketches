package main

import (
	"fmt"
	"time"
)

func accumulator(passing, completion chan int) {
	accumulate := <- passing
	fmt.Println("thread sleeping",accumulate)
	time.Sleep(time.Second)
	accumulate += 1
	completion <- 1
	passing <- accumulate
}

func main() {
	passing, completion := make(chan int), make(chan int)
	for i := 0; i < 5; i++ {
		go accumulator(passing, completion)
	}
	passing <- 0 
	for i := 0; i < 5; i++ {
		<- completion
	}
	accumulate := <- passing
	fmt.Println("final accumulate value:", accumulate)
}
