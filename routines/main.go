package main

import (
	"fmt"
)

type student struct {
	name       string
	rollNumber int
}

func main() {
	ch := make(chan int)
	go printNumbers(ch)
	<-ch
}

func printNumbers(ch chan int) {
	for i := range 10 {
		st := &student{name: "random", rollNumber: i}
		fmt.Println(*st)
	}

	ch <- 1
}
