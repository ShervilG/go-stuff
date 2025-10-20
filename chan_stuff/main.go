package main

import "sync"

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}

	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	for v := range ch {
		println(v)
	}

	wg.Done()
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go producer(ch)
	go consumer(ch, wg)
	wg.Wait()
}
