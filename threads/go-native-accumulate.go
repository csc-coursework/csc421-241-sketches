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
		// the write to the completion channel blocks
		// until there is an accumulator to read the channel
		<- completion
		fmt.Println("an accumulator thread wrote to the completion channel")
	}
	// note that accumlator reads passing before writing completion
	// therefore this read of passing will be the final write to passing
	accumulate := <- passing
	fmt.Println("final accumulate value:", accumulate)
}
