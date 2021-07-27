package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func exitOnSigterm() {
	signs := make(chan os.Signal, 1)

	signal.Notify(signs, syscall.SIGTERM)

	go func() {
		<-signs
		log.Println("Exiting from program in 1 second.")
		time.Sleep(time.Second)
		log.Println("Exit from program")
		os.Exit(0)
	}()
}

func main() {
	exitOnSigterm()
	for {
		time.Sleep(time.Second)
	}
}
