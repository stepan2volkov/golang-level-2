package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func calculateLong(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1e5; i += 1 {
	}
}

func main() {
	err := trace.Start(os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()
	goroutineNum := 3
	wg := &sync.WaitGroup{}
	wg.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		go calculateLong(wg)
		runtime.Gosched()
	}
	runtime.Gosched()
	time.Sleep(100 * time.Nanosecond)
}
