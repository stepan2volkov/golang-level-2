package main

import (
	"fmt"
	"sync"
)

var value int = 0

func main() {
	wg := &sync.WaitGroup{}
	timesToIncrement := 400

	wg.Add(timesToIncrement)
	for i := 0; i < timesToIncrement; i++ {
		go func() {
			defer wg.Done()
			value++
		}()
	}
	wg.Wait()
	fmt.Println(value)
}
