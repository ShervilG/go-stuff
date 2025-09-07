package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(5 * time.Second)
	startTime := time.Now()
	for currentTime := range tick {
		fmt.Println("Time passed !")
		if currentTime.Unix()-startTime.Unix() >= 14 {
			break
		}
	}

	fmt.Println("The end")
}
