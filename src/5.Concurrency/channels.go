package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Main func completed")
}

func work() {
	fmt.Println("Starting work method")
	time.Sleep(2 * time.Second)
	fmt.Println("Finishing work method")
}
