package main

/*
 * name: rendez-vous.go
 * last-update:
 *    4 oct 2025 -bjr: created
 *
 * two sample runs
 *
 * % make rendez-vous
 * go run rendez-vous.go
 * rv_a: a_g = 3, b_g = 5
 * rv_b: a_g = 3, b_g = 5
 * % make rendez-vous
 * go run rendez-vous.go
 * rv_b: a_g = 3, b_g = 5
 * rv_a: a_g = 3, b_g = 5
 *
 */

import (
	"fmt"
	"time"
)

// package-level variables
var a_g int = 0
var b_g int = 0

func rv_a(ch chan int) {
	a_g = 3
	<- ch
	// rendez-vous'ed
	fmt.Printf("rv_a: a_g = %d, b_g = %d\n", a_g, b_g)
}

func rv_b(ch chan int) {
	b_g = 5
	ch <- 1
	// rendez-vous'ed
	fmt.Printf("rv_b: a_g = %d, b_g = %d\n", a_g, b_g)
}

func main () {
	ch := make(chan int)
	go rv_a(ch)
	go rv_b(ch)
	time.Sleep(time.Second)
}

