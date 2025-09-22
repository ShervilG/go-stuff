package main

import (
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	startTime := time.Now()

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
	}()

	wg.Wait()
	elapsed := time.Since(startTime)
	println("All goroutines completed in", elapsed.Seconds(), "seconds")
}
