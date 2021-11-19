package main

import (
	"log"
	"sync"
	"time"
)

// 1. Напишите программу, которая запускает n потоков и дожидается завершения их всех
func doConcurrent(n int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, num int) {
			defer wg.Done()
			time.Sleep(time.Second)
			log.Printf("goroutine #%d done", num)
		}(wg, i)
	}
	wg.Wait()
}

func main() {
	doConcurrent(100)
}
