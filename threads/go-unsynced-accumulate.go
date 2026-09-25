//
// go-unsynced-accumulate.go
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

func accumulator() {
	var i = accumulate
	fmt.Println("thread sleeping on lock")
	time.Sleep(2*time.Second)
	accumulate = i+1
	fmt.Println("thread completing")
}

func main() {

	for i := 0; i < 5; i++ {
		go accumulator()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("final accumulate value:", accumulate)
}
