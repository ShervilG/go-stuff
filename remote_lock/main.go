package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var counter int

func main() {
	counter = 0

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			res := redisClient.SetNX("lock", "1", 5*time.Second)
			if res.Err() == nil && res.Val() {
				counter += 1
			} else {
				return
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
