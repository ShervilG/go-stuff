package main

import "sync"

var mut = &sync.Mutex{}

func incrementCounter(counter *int) {
	mut.Lock()
	*counter++
	mut.Unlock()
}

func main() {
	counter := 0
	wg := &sync.WaitGroup{}

	for range 10000 {
		wg.Add(1)
		go func() {
			incrementCounter(&counter)
			wg.Done()
		}()
	}

	wg.Wait()
	println("Final Counter:", counter)
}
