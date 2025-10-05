package main

/*
 * name: hello.go
 * last-update:
 *    4 oct 2025 -bjr: created
 *
 * two sample runs:
 *
 * % make hello
 * go run hello.go
 * ollhe
 * % make hello
 * go run hello.go
 * ohlel
 */

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Print(s)
}

func main() {
	hello := [5]string{"h","e","l","l","o"}
	i := 0 // := to define i, inferring type
	for i<len(hello) {    // or: i=0; i<5; i+=1 (no parens)
		go f(hello[i])
		i = i+1 // i updated, not defined
	}
	time.Sleep(time.Second)
	fmt.Println("")
}

